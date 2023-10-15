package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/execaus/exloggo"
	"simbir-go-api/constants"
	"simbir-go-api/models"
	"simbir-go-api/queries"
)

type AccountPostgres struct {
	db      *sql.DB
	queries *queries.Queries
}

func (r *AccountPostgres) GetList(start, count int32) ([]models.Account, error) {
	var accounts []models.Account

	accountRows, err := r.queries.GetAccounts(context.Background(), queries.GetAccountsParams{
		Offset: start,
		Limit:  count,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	for _, account := range accountRows {
		var roles []string

		if err = json.Unmarshal(account.Roles, &roles); err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}

		accounts = append(accounts, models.Account{
			Username: account.Username,
			Password: account.Password,
			Balance:  account.Balance,
			Roles:    roles,
		})
	}

	return accounts, nil
}

func (r *AccountPostgres) ReplaceUsername(username, newUsername string) error {
	if err := r.queries.ReplaceUsername(context.Background(), queries.ReplaceUsernameParams{
		Username:   newUsername,
		Username_2: username,
	}); err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *AccountPostgres) Update(username, newUsername, password string) error {
	if err := r.queries.UpdateAccount(context.Background(), queries.UpdateAccountParams{
		Username:   newUsername,
		Password:   password,
		Username_2: username,
	}); err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *AccountPostgres) BlockToken(token string) error {
	if err := r.queries.AppendTokenToBlackList(context.Background(), token); err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *AccountPostgres) IsContainBlackListToken(token string) (bool, error) {
	isContain, err := r.queries.IsContainBlackListToken(context.Background(), token)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isContain, err
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

func (r *AccountPostgres) create(username, password, role string, balance float64) (*queries.Account, error) {
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
		func(tx *queries.Queries) error {
			if role == constants.RoleUser {
				return nil
			}

			_, err := tx.AppendRoleAccount(context.Background(), queries.AppendRoleAccountParams{
				Account: username,
				Role:    constants.RoleUser,
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
