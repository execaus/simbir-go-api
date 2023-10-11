package service

import (
	"github.com/execaus/exloggo"
	"golang.org/x/crypto/bcrypt"
	"simbir-go-api/queries"
	"simbir-go-api/repository"
)

type AccountService struct {
	repo repository.Account
}

func (s *AccountService) IsExist(username string) (bool, error) {
	return s.repo.IsExist(username)
}

func (s *AccountService) SignUp(username, password string) (*queries.Account, error) {
	passwordHash, err := getHashPassword(password)
	if err != nil {
		return nil, err
	}

	account, err := s.repo.Create(username, passwordHash, false, 0)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func getHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		exloggo.Error(err.Error())
		return "", err
	}

	return string(hash), nil
}
