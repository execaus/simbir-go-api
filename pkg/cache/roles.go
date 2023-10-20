package cache

import (
	"errors"
	"github.com/execaus/exloggo"
	"simbir-go-api/types"
)

const (
	usernameNotFound = "username not found"
	rolesNotFound    = "roles not found"
)

type AccountRoleCache struct {
	Roles types.AccountRolesDictionary
}

func (c *AccountRoleCache) ReplaceRoles(username string, roles []string) error {
	_, ok := c.Roles.Load(username)
	if !ok {
		exloggo.Error(usernameNotFound)
		return errors.New(usernameNotFound)
	}

	c.Roles.Store(username, roles)

	return nil
}

func (c *AccountRoleCache) ReplaceUsername(username, newUsername string) error {
	currentRoles, ok := c.Roles.Load(username)
	if !ok {
		exloggo.Error(usernameNotFound)
		return errors.New(usernameNotFound)
	}

	c.Roles.Delete(username)
	c.Roles.Store(newUsername, currentRoles)
	return nil
}

func (c *AccountRoleCache) AppendRole(username string, newRole string) error {
	currentRoles, ok := c.Roles.Load(username)
	if !ok {
		c.Roles.Store(username, []string{newRole})
		return nil
	}

	for _, role := range currentRoles.([]string) {
		if role == newRole {
			return nil
		}
	}

	c.Roles.Store(username, append(currentRoles.([]string), newRole))

	return nil
}

func (c *AccountRoleCache) Load(m types.AccountRolesDictionary) {
	c.Roles = m
}

func (c *AccountRoleCache) GetRoles(username string) ([]string, error) {
	roles, ok := c.Roles.Load(username)
	if !ok {
		return nil, errors.New(rolesNotFound)
	}

	return roles.([]string), nil
}

func NewAccountRoleCache() *AccountRoleCache {
	var cache types.AccountRolesDictionary
	return &AccountRoleCache{Roles: cache}
}
