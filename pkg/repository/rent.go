package repository

import (
	"context"
	"github.com/execaus/exloggo"
	"simbir-go-api/queries"
)

type RentPostgres struct {
	queries *queries.Queries
}

func (r *RentPostgres) IsExist(id int32) (bool, error) {
	isExist, err := r.queries.IsRentExist(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isExist, err
}

func (r *RentPostgres) IsRemoved(id int32) (bool, error) {
	isRemoved, err := r.queries.IsRentRemoved(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isRemoved, err
}

func NewRentPostgres(queries *queries.Queries) *RentPostgres {
	return &RentPostgres{queries: queries}
}
