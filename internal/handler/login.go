package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"go.mod/internal/config"
	"go.mod/internal/model"
	"net/http"
	"strings"
	"time"
)

// LoginHandler — обработчик входа пользователя

func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Считываем данные из тела запроса
	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Проверка username и password в базе данных
	var role string
	err := h.db.Connection.QueryRow(
		"SELECT role FROM users WHERE username=$1 AND password=$2",
		creds.Username,
		creds.Password,
	).Scan(&role)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Создание JWT-токена
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &model.Claims{
		Username: creds.Username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Подпись токена секретным ключом
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWTKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	// Возвращаем токен в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// ========================== Проверка роли администратора ==========================

// IsAdmin — middleware для проверки, что пользователь является администратором.
func (h *Handlers) IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем токен из заголовка Authorization
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

		// Разбираем токен и проверяем его подпись
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JWTKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Проверка роли пользователя
		if claims.Role != "admin" {
			http.Error(w, "Forbidden: Admins only", http.StatusForbidden)
			return
		}

		// Всё хорошо — передаём управление следующему обработчику
		next(w, r)
	}
}
