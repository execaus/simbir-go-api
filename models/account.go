package models

type Account struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	IsAdmin  bool    `json:"isAdmin"`
	Balance  float64 `json:"balance"`
}

type SignUpInput struct {
	Username string `json:"username" binding:"required,excludesall= ,printascii"`
	Password string `json:"password" binding:"required,excludesall= ,printascii"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required,excludesall= ,printascii"`
	Password string `json:"password" binding:"required,excludesall= ,printascii"`
}
