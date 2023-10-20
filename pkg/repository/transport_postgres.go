package repository

import (
	"context"
	"github.com/execaus/exloggo"
	"simbir-go-api/models"
	"simbir-go-api/pkg/repository/sqlnt"
	"simbir-go-api/queries"
)

type TransportPostgres struct {
	queries *queries.Queries
}

func (r *TransportPostgres) IsExist(identifier string) (bool, error) {
	isExist, err := r.queries.IsExistTransport(context.Background(), identifier)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *TransportPostgres) Create(transport *models.Transport) (*models.Transport, error) {
	result, err := r.queries.CreateTransport(context.Background(), queries.CreateTransportParams{
		ID:          transport.Identifier,
		Owner:       transport.OwnerID,
		Type:        transport.TransportType,
		CanRanted:   transport.CanBeRented,
		Model:       transport.Model,
		Color:       transport.Color,
		Description: sqlnt.ToStringNull(transport.Description),
		Latitude:    transport.Latitude,
		Longitude:   transport.Longitude,
		MinutePrice: sqlnt.ToF64Null(transport.MinutePrice),
		DayPrice:    sqlnt.ToF64Null(transport.DayPrice),
	})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Transport{
		OwnerID:       result.Owner,
		CanBeRented:   result.CanRanted,
		TransportType: result.Type,
		Model:         result.Model,
		Color:         result.Color,
		Identifier:    result.ID,
		Description:   sqlnt.ToString(&result.Description),
		Latitude:      result.Latitude,
		Longitude:     result.Longitude,
		MinutePrice:   sqlnt.ToF64(&result.MinutePrice),
		DayPrice:      sqlnt.ToF64(&result.DayPrice),
	}, nil
}

func NewTransportPostgres(queries *queries.Queries) *TransportPostgres {
	return &TransportPostgres{queries: queries}
}
