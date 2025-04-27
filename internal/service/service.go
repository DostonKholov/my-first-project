package service

import "go.mod/internal/database"

type Service struct {
	database *database.Database
}

func NewService(db *database.Database) *Service {
	return &Service{
		database: db,
	}
}
