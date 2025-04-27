package model

import "github.com/golang-jwt/jwt/v5"

// Claims для JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// Структура для аутентификации
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
