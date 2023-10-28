package repository

import (
	"context"
	"github.com/execaus/exloggo"
	"simbir-go-api/queries"
)

type RentPostgres struct {
	queries *queries.Queries
}

func (r *RentPostgres) GetListFromUsername(username string) ([]queries.GetRentsFromUsernameRow, error) {
	rows, err := r.queries.GetRentsFromUsername(context.Background(), username)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return rows, nil
}

func (r *RentPostgres) GetListFromTransport(transportID string) ([]queries.GetRentsFromTransportIDRow, error) {
	rows, err := r.queries.GetRentsFromTransportID(context.Background(), transportID)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return rows, nil
}

func (r *RentPostgres) Get(id int32) (*queries.GetRentRow, error) {
	row, err := r.queries.GetRent(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &row, err
}

func (r *RentPostgres) IsRenter(id int32, username string) (bool, error) {
	isRenter, err := r.queries.IsRenter(context.Background(), queries.IsRenterParams{
		ID:      id,
		Account: username,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isRenter, err
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
