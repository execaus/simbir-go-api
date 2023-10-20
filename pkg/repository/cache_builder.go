package repository

import (
	"context"
	"github.com/execaus/exloggo"
	"simbir-go-api/queries"
	"simbir-go-api/types"
	"sync"
)

type CacheBuilderPostgres struct {
	queries *queries.Queries
}

func (c *CacheBuilderPostgres) CacheRoles() (types.AccountRolesDictionary, error) {
	var cacheRoles sync.Map

	result, err := c.queries.GetCacheRoles(context.Background())
	if err != nil {
		exloggo.Error(err.Error())
		return &cacheRoles, err
	}

	for _, accountRole := range result {
		currentRoles, ok := cacheRoles.Load(accountRole.Account)
		if !ok {
			cacheRoles.Store(accountRole.Account, []string{accountRole.Role})
		} else {
			cacheRoles.Store(accountRole.Account, append(currentRoles.([]string), accountRole.Role))
		}
	}

	return &cacheRoles, nil
}

func NewCacheBuilderPostgres(queries *queries.Queries) *CacheBuilderPostgres {
	return &CacheBuilderPostgres{queries: queries}
}
