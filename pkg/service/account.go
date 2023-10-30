package service

import (
	"encoding/json"
	"errors"
	"github.com/execaus/exloggo"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"simbir-go-api/constants"
	"simbir-go-api/models"
	"simbir-go-api/pkg/repository"
	"simbir-go-api/queries"
	"time"
)

const (
	invalidJwtMethod    = "invalid signing method"
	roleNotExist        = "role is not exist"
	tokenTTL            = 12 * time.Hour
	BalanceHesoyamValue = 250_000
)

type AccountService struct {
	cache repository.Role
	repo  repository.Account
	env   *models.Environment
}

func (s *AccountService) Hesoyam(id int32) (*models.Account, error) {
	account, err := s.repo.GetByID(id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	account.Balance += BalanceHesoyamValue

	if err = s.repo.Update(&models.Account{
		ID:       account.ID,
		Username: account.Username,
		Password: account.Password,
		Balance:  account.Balance,
	}); err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	roles, err := s.repo.GetRoles(id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Account{
		ID:        account.ID,
		Username:  account.Username,
		Password:  "",
		Balance:   account.Balance,
		Roles:     roles,
		IsDeleted: account.Deleted,
	}, nil
}

func (s *AccountService) IsRemovedByID(userID int32) (bool, error) {
	return s.repo.IsRemovedByID(userID)
}

func (s *AccountService) IsRemovedByUsername(username string) (bool, error) {
	return s.repo.IsRemovedByUsername(username)
}

func (s *AccountService) Remove(userID int32) error {
	return s.repo.RemoveAccount(userID)
}

func (s *AccountService) Create(username, password string, role string, balance float64) (*models.Account, error) {
	passwordHash, err := getPasswordHash(password)
	if err != nil {
		return nil, err
	}

	var account *queries.Account
	switch {
	case role == constants.RoleAdmin:
		account, err = s.repo.CreateAdmin(username, passwordHash, balance)
		if err != nil {
			return nil, err
		}
	case role == constants.RoleUser:
		account, err = s.repo.CreateUser(username, passwordHash, balance)
		if err != nil {
			return nil, err
		}
	default:
		exloggo.Error(roleNotExist)
		return nil, errors.New(roleNotExist)
	}

	account, err = s.repo.GetByID(account.ID)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	roles, err := s.repo.GetRoles(account.ID)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	for _, accountRole := range roles {
		if err = s.cache.AppendRole(account.ID, accountRole); err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
	}

	return &models.Account{
		ID:        account.ID,
		Username:  account.Username,
		Password:  account.Password,
		Balance:   account.Balance,
		Roles:     roles,
		IsDeleted: account.Deleted,
	}, nil
}

func (s *AccountService) GetList(start, count int32) ([]models.Account, error) {
	var accounts []models.Account

	reposAccounts, err := s.repo.GetList(start, count)
	if err != nil {
		return nil, err
	}

	for _, account := range reposAccounts {
		var roles []string

		if err = json.Unmarshal(account.Roles, &roles); err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}

		accounts = append(accounts, models.Account{
			ID:        account.ID,
			Username:  account.Username,
			Password:  account.Password,
			Balance:   account.Balance,
			Roles:     roles,
			IsDeleted: account.Deleted,
		})
	}

	return accounts, nil
}

func (s *AccountService) GetRoles(userID int32) ([]string, error) {
	return s.cache.GetRoles(userID)
}

func (s *AccountService) Update(updatedAccount *models.Account) (*models.Account, error) {
	passwordHash, err := getPasswordHash(updatedAccount.Username)
	if err != nil {
		return nil, err
	}

	updatedAccount.Password = passwordHash

	if err = s.repo.Update(updatedAccount); err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	if err = s.cache.ReplaceRoles(updatedAccount.ID, updatedAccount.Roles); err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return s.GetByID(updatedAccount.ID)
}

func (s *AccountService) IsValidToken(token string) (bool, error) {
	isContain, err := s.repo.IsContainBlackListToken(token)
	if err != nil {
		return false, err
	}

	return !isContain, err
}

func (s *AccountService) BlockToken(token string) error {
	return s.repo.BlockToken(token)
}

func (s *AccountService) GetByID(userID int32) (*models.Account, error) {
	account, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	roles, err := s.cache.GetRoles(userID)
	if err != nil {
		return nil, err
	}

	return &models.Account{
		ID:        account.ID,
		Username:  account.Username,
		Password:  "",
		Balance:   account.Balance,
		Roles:     roles,
		IsDeleted: account.Deleted,
	}, nil
}

func (s *AccountService) ParseToken(accessToken string) (int32, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.JWTTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(invalidJwtMethod)
		}

		return []byte(s.env.JWTSigningKey), nil
	})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*models.JWTTokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func (s *AccountService) Authorize(username, password string) (*models.Account, error) {
	account, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if err = comparePasswords(account.Password, password); err != nil {
		exloggo.Warning(err.Error())
		return nil, nil
	}

	roles, err := s.cache.GetRoles(account.ID)
	if err != nil {
		return nil, err
	}

	return &models.Account{
		ID:        account.ID,
		Username:  account.Username,
		Password:  account.Password,
		Roles:     roles,
		Balance:   account.Balance,
		IsDeleted: account.Deleted,
	}, nil
}

func (s *AccountService) GenerateToken(userID int32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	})

	signedString, err := token.SignedString([]byte(s.env.JWTSigningKey))
	if err != nil {
		exloggo.Error(err.Error())
		return "", err
	}

	return signedString, nil
}

func (s *AccountService) IsExistByID(userID int32) (bool, error) {
	return s.repo.IsExistByID(userID)
}

func (s *AccountService) IsExistByUsername(username string) (bool, error) {
	return s.repo.IsExistByUsername(username)
}

func (s *AccountService) SignUp(username, password string) (*models.Account, error) {
	passwordHash, err := getPasswordHash(password)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	account, err := s.repo.CreateUser(username, passwordHash, 0)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	if err = s.cache.AppendRole(account.ID, constants.RoleUser); err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	roles, err := s.repo.GetRoles(account.ID)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Account{
		ID:        account.ID,
		Username:  account.Username,
		Password:  "",
		Balance:   account.Balance,
		Roles:     roles,
		IsDeleted: account.Deleted,
	}, nil
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
