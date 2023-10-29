package cache

import (
	"errors"
	"github.com/execaus/exloggo"
	"simbir-go-api/types"
)

const (
	userIDNotFound = "user id not found"
	rolesNotFound  = "roles not found"
)

type AccountRoleCache struct {
	Roles types.AccountRolesDictionary
}

func (c *AccountRoleCache) ReplaceRoles(id int32, roles []string) error {
	_, ok := c.Roles.Load(id)
	if !ok {
		exloggo.Error(userIDNotFound)
		return errors.New(userIDNotFound)
	}

	c.Roles.Store(id, roles)

	return nil
}

func (c *AccountRoleCache) AppendRole(id int32, newRole string) error {
	currentRoles, ok := c.Roles.Load(id)
	if !ok {
		c.Roles.Store(id, []string{newRole})
		return nil
	}

	for _, role := range currentRoles.([]string) {
		if role == newRole {
			return nil
		}
	}

	c.Roles.Store(id, append(currentRoles.([]string), newRole))

	return nil
}

func (c *AccountRoleCache) Load(m types.AccountRolesDictionary) {
	c.Roles = m
}

func (c *AccountRoleCache) GetRoles(id int32) ([]string, error) {
	roles, ok := c.Roles.Load(id)
	if !ok {
		return nil, errors.New(rolesNotFound)
	}

	return roles.([]string), nil
}

func NewAccountRoleCache() *AccountRoleCache {
	var cache types.AccountRolesDictionary
	return &AccountRoleCache{Roles: cache}
}
