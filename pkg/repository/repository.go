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
	Update(username string, updatedAccount *models.Account) error
	GetList(start int32, count int32) ([]models.Account, error)
	RemoveAccount(username string) error
	IsRemoved(username string) (bool, error)
}

type Role interface {
	GetRoles(username string) ([]string, error)
	AppendRole(username string, role string) error
	ReplaceUsername(username, newUsername string) error
	ReplaceRoles(username string, roles []string) error
}

type CacheBuilder interface {
	CacheRoles() (types.AccountRolesDictionary, error)
}

type TransportRepository interface {
	Create(transport *models.Transport) (*models.Transport, error)
	IsExist(identifier string) (bool, error)
	Get(identifier string) (*models.Transport, error)
}

type Repository struct {
	Account
	CacheBuilder
	Transport TransportRepository
}

func NewRepository(queries *queries.Queries, db *sql.DB) *Repository {
	return &Repository{
		Account:      NewAccountPostgres(queries, db),
		CacheBuilder: NewCacheBuilderPostgres(queries),
		Transport:    NewTransportPostgres(queries),
	}
}
