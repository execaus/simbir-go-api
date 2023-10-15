package repository

import (
	"database/sql"
	"simbir-go-api/models"
	"simbir-go-api/queries"
	"simbir-go-api/types"
)

type Account interface {
	Role
	CreateUser(username, password string, balance float64) (*queries.Account, error)
	CreateAdmin(username, password string, balance float64) (*queries.Account, error)
	IsExist(username string) (bool, error)
	Get(username string) (*queries.Account, error)
	IsContainBlackListToken(token string) (bool, error)
	BlockToken(token string) error
	Update(username, newUsername, password string) error
	GetList(start int32, count int32) ([]models.Account, error)
}

type Role interface {
	GetRoles(username string) ([]string, error)
	AppendRole(username string, role string) error
	ReplaceUsername(username, newUsername string) error
}

type CacheBuilder interface {
	CacheRoles() (types.AccountRolesDictionary, error)
}

type Repository struct {
	Account
	CacheBuilder
}

func NewRepository(queries *queries.Queries, db *sql.DB) *Repository {
	return &Repository{
		Account:      NewAccountPostgres(queries, db),
		CacheBuilder: NewCacheBuilderPostgres(queries),
	}
}
