package service

import (
	"simbir-go-api/models"
	"simbir-go-api/pkg/cache"
	"simbir-go-api/pkg/repository"
)

type Account interface {
	GenerateToken(id int32) (string, error)
	ParseToken(token string) (int32, error)
	SignUp(username, password string) (*models.Account, error)
	IsExistByID(id int32) (bool, error)
	IsExistByUsername(username string) (bool, error)
	Authorize(username, password string) (*models.Account, error)
	GetByID(id int32) (*models.Account, error)
	IsValidToken(token string) (bool, error)
	BlockToken(token string) error
	Update(updatedAccount *models.Account) (*models.Account, error)
	GetRoles(id int32) ([]string, error)
	GetList(start, count int32) ([]models.Account, error)
	Create(username, password string, role string, balance float64) (*models.Account, error)
	Remove(id int32) error
	IsRemovedByID(id int32) (bool, error)
	IsRemovedByUsername(username string) (bool, error)
}

type Transport interface {
	Create(transport *models.Transport) (*models.Transport, error)
	IsExistByID(id int32) (bool, error)
	IsExistByIdentifier(identifier string) (bool, error)
	Get(id int32) (*models.Transport, error)
	IsOwner(id int32, userID int32) (bool, error)
	Update(transport *models.Transport) (*models.Transport, error)
	Remove(id int32) error
	IsRemoved(id int32) (bool, error)
	GetList(start, count int32, transportType string) ([]models.Transport, error)
	GetFromRadius(point *models.Point, radius float64, transportType string) ([]models.Transport, error)
	IsAccessRent(id int32) (bool, error)
}

type Rent interface {
	IsRemoved(id int32) (bool, error)
	IsExist(id int32) (bool, error)
	IsRenter(id int32, userID int32) (bool, error)
	Get(id int32) (*models.Rent, error)
	GetListFromUserID(id, start, count int32) ([]models.Rent, error)
	GetListFromTransportID(id, start, count int32) ([]models.Rent, error)
	TransportIsRented(id int32) (bool, error)
	Create(rent *models.Rent) (*models.Rent, error)
	End(id int32) (*models.Rent, error)
}

type Service struct {
	Account
	Transport
	Rent
}

func NewService(repos *repository.Repository, env *models.Environment, cache *cache.Cache) *Service {
	return &Service{
		Account:   NewAccountService(repos.Account, env, cache.Role),
		Transport: NewTransportService(repos.Transport),
		Rent:      NewRentService(repos.Rent),
	}
}
