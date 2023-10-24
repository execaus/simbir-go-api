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
	IsOwner(identifier, username string) (bool, error)
	Update(identifier string, transport *models.Transport) (*models.Transport, error)
	Remove(identifier string) error
	IsRemoved(identifier string) (bool, error)
	GetList(start, count int32) ([]queries.Transport, error)
	GetListOnlyType(start, count int32, transportType string) ([]queries.Transport, error)
	GetFromRadiusAll(point *models.Point, radius float64, transportType string) ([]queries.Transport, error)
	GetFromRadiusOnlyType(point *models.Point, radius float64, transportType string) ([]queries.Transport, error)
}

type Rent interface {
	IsRemoved(id int32) (bool, error)
	IsExist(id int32) (bool, error)
	IsRenter(id int32, username string) (bool, error)
	Get(id int32) (*queries.GetRentRow, error)
}

type Repository struct {
	Account
	CacheBuilder
	Transport TransportRepository
	Rent
}

func NewRepository(queries *queries.Queries, db *sql.DB) *Repository {
	return &Repository{
		Account:      NewAccountPostgres(queries, db),
		CacheBuilder: NewCacheBuilderPostgres(queries),
		Transport:    NewTransportPostgres(queries),
		Rent:         NewRentPostgres(queries),
	}
}
