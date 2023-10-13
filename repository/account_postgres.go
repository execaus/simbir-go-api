package repository

import (
	"context"
	"github.com/execaus/exloggo"
	"simbir-go-api/queries"
)

type AccountPostgres struct {
	db *queries.Queries
}

func (r *AccountPostgres) Get(username string) (*queries.Account, error) {
	account, err := r.db.GetAccount(context.Background(), username)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func (r *AccountPostgres) IsExist(username string) (bool, error) {
	isExist, err := r.db.IsAccountExist(context.Background(), username)

	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *AccountPostgres) Create(username, password string, isAdmin bool, balance float64) (*queries.Account, error) {
	account, err := r.db.CreateAccount(context.Background(), queries.CreateAccountParams{
		Username: username,
		Password: password,
		IsAdmin:  isAdmin,
		Balance:  balance,
	})

	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func NewAccountPostgres(db *queries.Queries) *AccountPostgres {
	return &AccountPostgres{db: db}
}
