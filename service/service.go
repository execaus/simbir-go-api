package service

import (
	"simbir-go-api/queries"
	"simbir-go-api/repository"
)

type Account interface {
	SignUp(username, password string) (*queries.Account, error)
	IsExist(username string) (bool, error)
}

type Service struct {
	Account
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
	}
}
