package repository

import (
	"context"
	"github.com/execaus/exloggo"
	"simbir-go-api/pkg/repository/sqlnt"
	"simbir-go-api/queries"
	"time"
)

type RentPostgres struct {
	queries *queries.Queries
}

func (r *RentPostgres) End(id int32, timeEnd *time.Time) error {
	if err := r.queries.EndRent(context.Background(), queries.EndRentParams{
		TimeEnd: sqlnt.ToTimeNull(timeEnd),
		ID:      id,
	}); err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *RentPostgres) Create(username, transportID string, timeStart time.Time, timeEnd *time.Time, priceUnit float64, rentType string) (*queries.Rent, error) {
	rent, err := r.queries.CreateRent(context.Background(), queries.CreateRentParams{
		Account:   username,
		Transport: transportID,
		TimeStart: timeStart,
		TimeEnd:   sqlnt.ToTimeNull(timeEnd),
		PriceUnit: priceUnit,
		PriceType: rentType,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &rent, nil
}

func (r *RentPostgres) IsExistCurrentRented(transportID string) (bool, error) {
	isExist, err := r.queries.IsExistCurrentRent(context.Background(), transportID)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isExist, nil
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
