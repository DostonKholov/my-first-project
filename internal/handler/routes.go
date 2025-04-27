package handler

import (
	"github.com/gorilla/mux"
	"go.mod/internal/database"
	"go.mod/internal/service"
	"net/http"
)

// Handlers — структура для всех обработчиков
type Handlers struct {
	db      *database.Database // Работа с базой данных
	service *service.Service   // Логика приложения (если используется)
}

// NewHandler — конструктор нового экземпляра Handlers
func NewHandler(s *service.Service, db *database.Database) *Handlers {
	return &Handlers{
		service: s,
		db:      db,
	}
}

// InitRoutes — инициализация всех маршрутов (роутов) приложения
func (h *Handlers) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Авторизация
	router.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost, http.MethodOptions)

	// Открытые маршруты для пользователей
	router.HandleFunc("/employee", h.JWTMiddleware(h.GetEmployee)).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/employees", h.JWTMiddleware(h.GetAllEmployees)).Methods(http.MethodGet, http.MethodOptions)

	// Защищённые маршруты (только для админов)
	router.HandleFunc("/add_employee", h.JWTMiddleware(h.IsAdmin(h.CreateEmployee))).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete_employee", h.JWTMiddleware(h.IsAdmin(h.DeleteEmployee))).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/update_employee", h.JWTMiddleware(h.IsAdmin(h.UpdateEmployee))).Methods(http.MethodPost, http.MethodOptions)

	// Защищённые маршруты (только для админов)
	router.HandleFunc("/user", h.JWTMiddleware(h.IsAdmin(h.GetUser))).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users", h.JWTMiddleware(h.IsAdmin(h.GetAllUsers))).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/add_user", h.JWTMiddleware(h.IsAdmin(h.CreateUser))).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete_user", h.JWTMiddleware(h.IsAdmin(h.DeleteUser))).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/update_user", h.JWTMiddleware(h.IsAdmin(h.UpdateUser))).Methods(http.MethodPut, http.MethodOptions)

	return router
}
