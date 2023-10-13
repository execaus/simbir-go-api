package service

import (
	"github.com/execaus/exloggo"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"simbir-go-api/models"
	"simbir-go-api/queries"
	"simbir-go-api/repository"
	"time"
)

type AccountService struct {
	repo repository.Account
	env  *models.Environment
}

const tokenTTL = 12 * time.Hour

func (s *AccountService) Authorize(username, password string) (*models.Account, error) {
	account, err := s.repo.Get(username)
	if err != nil {
		return nil, err
	}

	if err = comparePasswords(account.Password, password); err != nil {
		exloggo.Warning(err.Error())
		return nil, nil
	}

	return &models.Account{
		Username: account.Username,
		Password: account.Password,
		IsAdmin:  account.IsAdmin,
		Balance:  account.Balance,
	}, nil
}

func (s *AccountService) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTTokenClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	signedString, err := token.SignedString([]byte(s.env.JWTSigningKey))
	if err != nil {
		exloggo.Error(err.Error())
		return "", err
	}

	return signedString, nil
}

func (s *AccountService) IsExist(username string) (bool, error) {
	return s.repo.IsExist(username)
}

func (s *AccountService) SignUp(username, password string) (*queries.Account, error) {
	passwordHash, err := getPasswordHash(password)
	if err != nil {
		return nil, err
	}

	account, err := s.repo.Create(username, passwordHash, false, 0)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func NewAccountService(repo repository.Account, env *models.Environment) *AccountService {
	return &AccountService{repo: repo, env: env}
}

func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		exloggo.Error(err.Error())
		return "", err
	}

	return string(hash), nil
}

func comparePasswords(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
