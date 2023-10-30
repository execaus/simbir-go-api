package repository

import (
	"database/sql"
	"simbir-go-api/models"
	"simbir-go-api/queries"
	"simbir-go-api/types"
	"time"
)

type Account interface {
	Role
	CreateUser(username, password string, balance float64) (*queries.Account, error)
	CreateAdmin(username, password string, balance float64) (*queries.Account, error)
	IsExistByID(id int32) (bool, error)
	IsExistByUsername(username string) (bool, error)
	GetByID(id int32) (*queries.Account, error)
	GetByUsername(username string) (*queries.Account, error)
	IsContainBlackListToken(token string) (bool, error)
	BlockToken(token string) error
	Update(updatedAccount *models.Account) error
	GetList(start, count int32) ([]queries.GetExistAccountsRow, error)
	RemoveAccount(id int32) error
	IsRemovedByID(id int32) (bool, error)
	IsRemovedByUsername(username string) (bool, error)
}

type Role interface {
	GetRoles(id int32) ([]string, error)
	AppendRole(id int32, role string) error
	ReplaceRoles(id int32, roles []string) error
}

type CacheBuilder interface {
	CacheRoles() (types.AccountRolesDictionary, error)
}

type TransportRepository interface {
	Create(transport *models.Transport) (*models.Transport, error)
	IsExistByID(id int32) (bool, error)
	IsExistByIdentifier(identifier string) (bool, error)
	Get(id int32) (*models.Transport, error)
	IsOwner(id, userID int32) (bool, error)
	Update(transport *models.Transport) (*models.Transport, error)
	Remove(id int32) error
	IsRemoved(id int32) (bool, error)
	GetList(start, count int32) ([]queries.Transport, error)
	GetListOnlyType(start, count int32, transportType string) ([]queries.Transport, error)
	GetFromRadiusAll(point *models.Point, radiusForMile float64) ([]queries.Transport, error)
	GetFromRadiusOnlyType(point *models.Point, radiusForMile float64, transportType string) ([]queries.Transport, error)
}

type Rent interface {
	IsRemoved(id int32) (bool, error)
	IsExist(id int32) (bool, error)
	IsRenter(id int32, userID int32) (bool, error)
	Get(id int32) (*queries.Rent, error)
	GetListFromUserID(id, start, count int32) ([]queries.Rent, error)
	GetListFromTransportID(id, start, count int32) ([]queries.Rent, error)
	IsExistCurrentRented(id int32) (bool, error)
	Create(rent *models.Rent) (*queries.Rent, error)
	End(id int32, timeEnd *time.Time) error
	Update(rent *models.Rent) (*queries.Rent, error)
	Remove(id int32) error
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
