package models

type Account struct {
	Username string
	Password string
	IsAdmin  bool
	Balance  float64
}

type SignUpInput struct {
	Username string `json:"username" binding:"required,excludesall= ,printascii"`
	Password string `json:"password" binding:"required,excludesall= ,printascii"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required,excludesall= ,printascii"`
	Password string `json:"password" binding:"required,excludesall= ,printascii"`
}
