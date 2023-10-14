package cache

import (
	"simbir-go-api/pkg/repository"
	"simbir-go-api/types"
)

type Role interface {
	Load(types.AccountRolesDictionary)
	repository.Role
}

type Cache struct {
	Role
}

func NewCache() *Cache {
	return &Cache{Role: NewAccountRoleCache()}
}
