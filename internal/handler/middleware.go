package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mod/internal/config"
	"go.mod/internal/model"
	"net/http"
	"strings"
)

// JWTMiddleware — промежуточный обработчик для проверки JWT-токена
func (h *Handlers) JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем токен из заголовка Authorization
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if tokenStr == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Структура для хранения данных токена
		claims := &model.Claims{}

		// Парсим и проверяем токен
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JWTKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Сохраняем имя пользователя и роль в заголовок запроса (опционально)
		r.Header.Set("X-User", claims.Username)
		r.Header.Set("X-Role", claims.Role)

		// Вызываем следующий обработчик
		next(w, r)
	}
}
