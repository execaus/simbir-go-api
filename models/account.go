package models

import "simbir-go-api/constants"

type Account struct {
	Username string
	Password string
	Balance  float64
	Roles    []string
}

func (a *Account) IsAdmin() bool {
	for _, role := range a.Roles {
		if role == constants.RoleAdmin {
			return true
		}
	}
	return false
}

type SignUpInput struct {
	Username string `json:"username" binding:"required,excludesall= ,printascii"`
	Password string `json:"password" binding:"required,excludesall= ,printascii"`
}

type SignUpOutput struct {
	Token string `json:"token"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required,excludesall= ,printascii"`
	Password string `json:"password" binding:"required,excludesall= ,printascii"`
}

type SignInOutput struct {
	Token string `json:"token"`
}

type GetAccountOutput struct {
	Username string  `json:"username"`
	IsAdmin  bool    `json:"isAdmin"`
	Balance  float64 `json:"balance"`
}

type UpdateAccountInput struct {
	Username string `json:"username" binding:"required,excludesall= ,printascii"`
	Password string `json:"password" binding:"required,excludesall= ,printascii"`
}

type UpdateAccountOutput struct {
	Token string `json:"token"`
}
