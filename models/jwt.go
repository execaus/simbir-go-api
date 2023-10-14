package models

import "github.com/golang-jwt/jwt/v5"

type JWTTokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
