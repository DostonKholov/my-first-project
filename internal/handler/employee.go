package handler

import (
	"encoding/json"
	"fmt"
	"go.mod/internal/model"
	"net/http"
	"strconv"
)

// Получить одного сотрудника по ID
func (h *Handlers) GetEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Параметр 'id' отсутствует", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		http.Error(w, "Некорректный параметр 'id'", http.StatusBadRequest)
		return
	}

	employee, err := h.db.GetEmployeeByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		return
	}
}

// Получить всех сотрудников
func (h *Handlers) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	employees, err := h.db.GetAllEmployees()
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		return
	}
}

// Создать нового сотрудника
func (h *Handlers) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	if err := h.db.CreateEmployee(employee); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка создания сотрудника: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Сотрудник успешно создан"})
	if err != nil {
		return
	}
}

// Удалить сотрудника
func (h *Handlers) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Параметр 'id' отсутствует", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		http.Error(w, "Некорректный параметр 'id'", http.StatusBadRequest)
		return
	}

	if err := h.db.DeleteEmployee(id); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка удаления сотрудника: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Сотрудник успешно удалён"})
	if err != nil {
		return
	}
}

// Обновить данные сотрудника
func (h *Handlers) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Параметр 'id' отсутствует", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		http.Error(w, "Некорректный параметр 'id'", http.StatusBadRequest)
		return
	}

	if err := h.db.UpdateEmployee(id, employee); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка обновления сотрудника: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Сотрудник успешно обновлён"})
	if err != nil {
		return
	}
}
