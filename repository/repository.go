package repository

import (
	"simbir-go-api/queries"
)

type Account interface {
	Create(username, password string, isAdmin bool, balance float64) (*queries.Account, error)
	IsExist(username string) (bool, error)
}

type Repository struct {
	Account
}

func NewRepository(db *queries.Queries) *Repository {
	return &Repository{
		Account: NewAccountPostgres(db),
	}
}
