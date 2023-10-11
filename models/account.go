package models

type SignUpInput struct {
	Username string `json:"firstName" binding:"required,excludesall= ,printascii"`
	Password string `json:"password" binding:"required,excludesall= ,printascii"`
}
