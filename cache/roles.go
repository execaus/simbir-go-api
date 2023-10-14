package cache

import "simbir-go-api/types"

type AccountRoleCache struct {
	Roles types.AccountRolesDictionary
}

func (c *AccountRoleCache) AppendRole(username string, newRole string) error {
	if c.Roles[username] == nil {
		c.Roles[username] = []string{newRole}
	} else {
		for _, role := range c.Roles[username] {
			if role == newRole {
				return nil
			}
		}
		c.Roles[username] = append(c.Roles[username], newRole)
	}
	return nil
}

func (c *AccountRoleCache) Load(m types.AccountRolesDictionary) {
	c.Roles = m
}

func (c *AccountRoleCache) GetRoles(username string) ([]string, error) {
	return c.Roles[username], nil
}

func NewAccountRoleCache() *AccountRoleCache {
	return &AccountRoleCache{Roles: types.AccountRolesDictionary{}}
}
