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
	Update(username string, newUsername string, password string) (string, error)
	GetRoles(username string) ([]string, error)
	GetList(start, count int32) ([]models.Account, error)
	Create(username, password string, role string, balance float64) (*models.Account, error)
}

type Service struct {
	Account
}

func NewService(repos *repository.Repository, env *models.Environment, cache *cache.Cache) *Service {
	return &Service{
		Account: NewAccountService(repos.Account, env, cache.Role),
	}
}
