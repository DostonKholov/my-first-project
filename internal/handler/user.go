package handler

import (
	"encoding/json"
	"fmt"
	"go.mod/internal/model"
	"net/http"
	"strconv"
)

// Пользователи

// Получение одного пользователя по ID
func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	idString := r.URL.Query()["id"]
	if len(idString) == 0 {
		http.Error(w, "Параметр 'id' отсутствует", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idString[0], 10, 64)
	if err != nil {
		http.Error(w, "Некорректный параметр 'id'", http.StatusBadRequest)
		return
	}

	user, err := h.db.GetUserByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка получения пользователя: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

// Получение всех пользователей
func (h *Handlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	users, err := h.db.GetAllUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка получения пользователей: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}

// Создание нового пользователя
func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	if err := h.db.CreateUser(user); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка создания пользователя: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Пользователь успешно создан"})
	if err != nil {
		return
	}
}

// Удаление пользователя
func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	idString := r.URL.Query()["id"]
	if len(idString) == 0 {
		http.Error(w, "Параметр 'id' отсутствует", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idString[0], 10, 64)
	if err != nil {
		http.Error(w, "Некорректный параметр 'id'", http.StatusBadRequest)
		return
	}

	if err := h.db.DeleteUser(id); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка удаления пользователя: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Пользователь успешно удалён"})
	if err != nil {
		return
	}
}

// Обновление пользователя
func (h *Handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	idString := r.URL.Query()["id"]
	if len(idString) == 0 {
		http.Error(w, "Параметр 'id' отсутствует", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idString[0], 10, 64)
	if err != nil {
		http.Error(w, "Некорректный параметр 'id'", http.StatusBadRequest)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	if err := h.db.UpdateUser(id, user); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка обновления пользователя: %v", err), http.StatusInternalServerError)
		return
	}

	// Ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Пользователь успешно обновлён"})
	if err != nil {
		return
	}
}
