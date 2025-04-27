package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mod/internal/database"
	"go.mod/internal/handler"
	"go.mod/internal/server"
	"go.mod/internal/service"
	"go.mod/pkg"
	_ "gorm.io/driver/postgres"
	"log"
)

func main() {
	// 1. Загружаем конфиг из config.yaml
	if err := pkg.InitConfig(); err != nil {
		log.Fatal("Ошибка загрузки конфига:", err)
	}

	// 2. Загружаем переменные окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	// 3. Подключение к базе данных
	connection := database.NewConnectPostgres()

	// 4. Создание объекта базы данных
	db := database.NewDatabase(connection)

	// 5. Создание сервисов (бизнес-логики)
	services := service.NewService(db)

	// 6. Создание обработчиков
	handler := handler.NewHandler(services, db) // исправил здесь

	// 7. Создание и запуск сервера
	app := new(server.Server)
	if err := app.ServerRun(handler.InitRoutes(), "8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
