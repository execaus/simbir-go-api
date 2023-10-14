package service

import (
	"errors"
	"github.com/execaus/exloggo"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"simbir-go-api/constants"
	"simbir-go-api/models"
	"simbir-go-api/queries"
	"simbir-go-api/repository"
	"time"
)

type AccountService struct {
	cache repository.Role
	repo  repository.Account
	env   *models.Environment
}

const (
	invalidJwtMethod = "invalid signing method"
	tokenTTL         = 12 * time.Hour
)

func (s *AccountService) GetByUsername(username string) (*models.Account, error) {
	account, err := s.repo.Get(username)
	if err != nil {
		return nil, err
	}

	roles, err := s.cache.GetRoles(username)
	if err != nil {
		return nil, err
	}

	return &models.Account{
		Username: account.Username,
		Password: "",
		Balance:  account.Balance,
		Roles:    roles,
	}, nil
}

func (s *AccountService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.JWTTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(invalidJwtMethod)
		}

		return []byte(s.env.JWTSigningKey), nil
	})
	if err != nil {
		return "", nil
	}

	claims, ok := token.Claims.(*models.JWTTokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Username, nil
}

func (s *AccountService) Authorize(username, password string) (*models.Account, error) {
	account, err := s.repo.Get(username)
	if err != nil {
		return nil, err
	}

	if err = comparePasswords(account.Password, password); err != nil {
		exloggo.Warning(err.Error())
		return nil, nil
	}

	roles, err := s.cache.GetRoles(username)
	if err != nil {
		return nil, err
	}

	return &models.Account{
		Username: account.Username,
		Password: account.Password,
		Roles:    roles,
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

	account, err := s.repo.CreateUser(username, passwordHash, 0)
	if err != nil {
		return nil, err
	}

	if err = s.cache.AppendRole(account.Username, constants.RoleUser); err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return account, nil
}

func NewAccountService(repo repository.Account, env *models.Environment, cache repository.Role) *AccountService {
	return &AccountService{repo: repo, env: env, cache: cache}
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
