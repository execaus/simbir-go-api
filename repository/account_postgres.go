package repository

import (
	"context"
	"database/sql"
	"github.com/execaus/exloggo"
	"simbir-go-api/constants"
	"simbir-go-api/queries"
)

type AccountPostgres struct {
	db      *sql.DB
	queries *queries.Queries
}

func (r *AccountPostgres) AppendRole(username string, role string) error {
	_, err := r.queries.AppendRoleAccount(context.Background(), queries.AppendRoleAccountParams{
		Account: username,
		Role:    role,
	})
	if err != nil {
		exloggo.Error(err.Error())
	}
	return err
}

func (r *AccountPostgres) GetRoles(username string) ([]string, error) {
	roles, err := r.queries.GetAccountRoles(context.Background(), username)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return roles, nil
}

func (r *AccountPostgres) Get(username string) (*queries.Account, error) {
	account, err := r.queries.GetAccount(context.Background(), username)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func (r *AccountPostgres) IsExist(username string) (bool, error) {
	isExist, err := r.queries.IsAccountExist(context.Background(), username)

	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isExist, nil
}

func (r *AccountPostgres) CreateUser(username, password string, balance float64) (*queries.Account, error) {
	return r.create(username, password, constants.RoleUser, balance)
}

func (r *AccountPostgres) CreateAdmin(username, password string, balance float64) (*queries.Account, error) {
	return r.create(username, password, constants.RoleAdmin, balance)
}

func (r *AccountPostgres) create(username, password string, role string, balance float64) (*queries.Account, error) {
	var account *queries.Account

	if err := r.ExecuteWithTransaction([]TXQuery{
		func(tx *queries.Queries) error {
			dbAccount, err := tx.CreateAccount(context.Background(), queries.CreateAccountParams{
				Username: username,
				Password: password,
				Balance:  balance,
			})
			if err != nil {
				exloggo.Error(err.Error())
				return err
			}

			account = &dbAccount

			return nil
		},
		func(tx *queries.Queries) error {
			_, err := tx.AppendRoleAccount(context.Background(), queries.AppendRoleAccountParams{
				Account: username,
				Role:    role,
			})
			if err != nil {
				exloggo.Error(err.Error())
			}
			return err
		},
	}); err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return account, nil
}

func NewAccountPostgres(queries *queries.Queries, db *sql.DB) *AccountPostgres {
	return &AccountPostgres{queries: queries, db: db}
}
