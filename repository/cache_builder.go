package repository

import (
	"context"
	"github.com/execaus/exloggo"
	"simbir-go-api/queries"
	"simbir-go-api/types"
)

type CacheBuilderPostgres struct {
	queries *queries.Queries
}

func (c *CacheBuilderPostgres) CacheRoles() (types.AccountRolesDictionary, error) {
	cacheRoles := make(types.AccountRolesDictionary)

	result, err := c.queries.GetCacheRoles(context.Background())
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	for _, accountRole := range result {
		if cacheRoles[accountRole.Account] == nil {
			cacheRoles[accountRole.Account] = []string{accountRole.Role}
		} else {
			cacheRoles[accountRole.Account] = append(cacheRoles[accountRole.Account], accountRole.Role)
		}
	}

	return cacheRoles, nil
}

func NewCacheBuilderPostgres(queries *queries.Queries) *CacheBuilderPostgres {
	return &CacheBuilderPostgres{queries: queries}
}
