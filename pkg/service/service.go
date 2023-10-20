package service

import (
	"simbir-go-api/models"
	"simbir-go-api/pkg/cache"
	"simbir-go-api/pkg/repository"
	"simbir-go-api/queries"
)

type Account interface {
	GenerateToken(username string) (string, error)
	ParseToken(token string) (string, error)
	SignUp(username, password string) (*queries.Account, error)
	IsExist(username string) (bool, error)
	Authorize(username, password string) (*models.Account, error)
	GetByUsername(username string) (*models.Account, error)
	IsValidToken(token string) (bool, error)
	BlockToken(token string) error
	Update(username string, updatedAccount *models.Account) (*models.Account, error)
	GetRoles(username string) ([]string, error)
	GetList(start, count int32) ([]models.Account, error)
	Create(username, password string, role string, balance float64) (*models.Account, error)
	Remove(username string) error
	IsRemoved(username string) (bool, error)
}

type Transport interface {
	Create(transport *models.Transport) (*models.Transport, error)
	IsExist(identifier string) (bool, error)
	Get(identifier string) (*models.Transport, error)
	IsOwner(identifier, username string) (bool, error)
	Update(identifier string, transport *models.Transport) (*models.Transport, error)
}

type Service struct {
	Account
	Transport
}

func NewService(repos *repository.Repository, env *models.Environment, cache *cache.Cache) *Service {
	return &Service{
		Account:   NewAccountService(repos.Account, env, cache.Role),
		Transport: NewTransportService(repos.Transport),
	}
}
