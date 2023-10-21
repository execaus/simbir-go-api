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

func (r *TransportPostgres) IsRemoved(identifier string) (bool, error) {
	isRemoved, err := r.queries.IsTransportRemoved(context.Background(), identifier)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isRemoved, nil
}

func (r *TransportPostgres) Remove(identifier string) error {
	if err := r.queries.RemoveTransport(context.Background(), identifier); err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *TransportPostgres) Update(identifier string, transport *models.Transport) (*models.Transport, error) {
	reposTransport, err := r.queries.UpdateTransport(context.Background(), queries.UpdateTransportParams{
		ID:          transport.Identifier,
		CanRanted:   transport.CanBeRented,
		Model:       transport.Model,
		Color:       transport.Color,
		Description: sqlnt.ToStringNull(transport.Description),
		Latitude:    transport.Latitude,
		Longitude:   transport.Longitude,
		MinutePrice: sqlnt.ToF64Null(transport.MinutePrice),
		DayPrice:    sqlnt.ToF64Null(transport.DayPrice),
		ID_2:        identifier,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Transport{
		OwnerID:       reposTransport.Owner,
		CanBeRented:   reposTransport.CanRanted,
		TransportType: reposTransport.Type,
		Model:         reposTransport.Model,
		Color:         reposTransport.Color,
		Identifier:    reposTransport.ID,
		Description:   sqlnt.ToString(&reposTransport.Description),
		Latitude:      reposTransport.Latitude,
		Longitude:     reposTransport.Longitude,
		MinutePrice:   sqlnt.ToF64(&reposTransport.MinutePrice),
		DayPrice:      sqlnt.ToF64(&reposTransport.MinutePrice),
		IsDeleted:     reposTransport.Deleted,
	}, nil
}

func (r *TransportPostgres) IsOwner(identifier, username string) (bool, error) {
	isOwner, err := r.queries.IsTransportOwner(context.Background(), queries.IsTransportOwnerParams{
		ID:    identifier,
		Owner: username,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isOwner, nil
}

func (r *TransportPostgres) Get(identifier string) (*models.Transport, error) {
	result, err := r.queries.GetTransport(context.Background(), identifier)
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
		IsDeleted:     result.Deleted,
	}, nil
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
		IsDeleted:     result.Deleted,
	}, nil
}

func NewTransportPostgres(queries *queries.Queries) *TransportPostgres {
	return &TransportPostgres{queries: queries}
}
