package service

import (
	"simbir-go-api/cache"
	"simbir-go-api/models"
	"simbir-go-api/queries"
	"simbir-go-api/repository"
)

type Account interface {
	GenerateToken(username string) (string, error)
	ParseToken(token string) (string, error)
	SignUp(username, password string) (*queries.Account, error)
	IsExist(username string) (bool, error)
	Authorize(username, password string) (*models.Account, error)
	GetByUsername(username string) (*models.Account, error)
}

type Service struct {
	Account
}

func NewService(repos *repository.Repository, env *models.Environment, cache *cache.Cache) *Service {
	return &Service{
		Account: NewAccountService(repos.Account, env, cache.Role),
	}
}
