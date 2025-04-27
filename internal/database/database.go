package database

import (
	"database/sql"
)

// Структура Database будет хранить подключение к базе
type Database struct {
	Connection *sql.DB
}

// Конструктор нового объекта базы данных
func NewDatabase(db *sql.DB) *Database {
	return &Database{
		Connection: db,
	}
}

//import (
//	"database/sql"
//)
//
//type Database struct {
//	connection *sql.DB
//}
//
//func NewDatabase(db *sql.DB) *Database {
//	return &Database{
//		connection: db,
//	}
//}
