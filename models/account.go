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

type AdminGetAccountsInput struct {
	Start int32 `form:"start" binding:"min=0"`
	Count int32 `form:"count" binding:"min=1"`
}

type AdminGetAccountsOutput struct {
	Accounts []*GetAccountOutput `json:"accounts"`
}

type AdminGetAccountInput struct {
	Username string `json:"username" binding:"required,excludesall= ,printascii"`
}

type AdminGetAccountOutput struct {
	Account GetAccountOutput `json:"account"`
}

type AdminCreateAccountInput struct {
	Username string   `json:"username" binding:"required,excludesall= ,printascii"`
	Password string   `json:"password" binding:"required,excludesall= ,printascii"`
	IsAdmin  bool     `json:"isAdmin" binding:"required"`
	Balance  *float64 `json:"balance" binding:"required,min=0"`
}

type AdminCreateAccountOutput struct {
	Account GetAccountOutput `json:"account"`
}
