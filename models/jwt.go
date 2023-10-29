package models

import "github.com/golang-jwt/jwt/v5"

type JWTTokenClaims struct {
	UserID int32 `json:"userID"`
	jwt.RegisteredClaims
}
