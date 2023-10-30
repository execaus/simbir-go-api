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

func (r *TransportPostgres) GetFromRadiusAll(point *models.Point, radiusForMile float64) ([]queries.Transport, error) {
	transports, err := r.queries.GetTransportsFromRadiusAll(context.Background(), queries.GetTransportsFromRadiusAllParams{
		Point:     point.Longitude,
		Point_2:   point.Latitude,
		Longitude: radiusForMile,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return transports, nil
}

func (r *TransportPostgres) GetFromRadiusOnlyType(point *models.Point, radiusForMile float64, transportType string) ([]queries.Transport, error) {
	transports, err :=
		r.queries.GetTransportsFromRadiusOnlyType(context.Background(), queries.GetTransportsFromRadiusOnlyTypeParams{
			Type:      transportType,
			Point:     point.Longitude,
			Point_2:   point.Latitude,
			Longitude: radiusForMile,
		})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return transports, nil
}

func (r *TransportPostgres) GetList(start, count int32) ([]queries.Transport, error) {
	transports, err := r.queries.GetTransports(context.Background(), queries.GetTransportsParams{
		Offset: start,
		Limit:  count,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return transports, nil
}

func (r *TransportPostgres) GetListOnlyType(start, count int32, transportType string) ([]queries.Transport, error) {
	transports, err := r.queries.GetTransportsOnlyType(context.Background(), queries.GetTransportsOnlyTypeParams{
		Offset: start,
		Limit:  count,
		Type:   transportType,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return transports, nil
}

func (r *TransportPostgres) IsRemoved(id int32) (bool, error) {
	isRemoved, err := r.queries.IsTransportRemoved(context.Background(), id)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isRemoved, nil
}

func (r *TransportPostgres) Remove(id int32) error {
	if err := r.queries.RemoveTransport(context.Background(), id); err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *TransportPostgres) Update(transport *models.Transport) (*models.Transport, error) {
	reposTransport, err := r.queries.UpdateTransport(context.Background(), queries.UpdateTransportParams{
		Identifier:  transport.Identifier,
		CanRented:   transport.CanBeRented,
		Model:       transport.Model,
		Color:       transport.Color,
		Description: sqlnt.ToStringNull(transport.Description),
		Latitude:    transport.Latitude,
		Longitude:   transport.Longitude,
		MinutePrice: sqlnt.ToF64Null(transport.MinutePrice),
		DayPrice:    sqlnt.ToF64Null(transport.DayPrice),
		ID:          transport.ID,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Transport{
		ID:            transport.ID,
		OwnerID:       reposTransport.Owner,
		CanBeRented:   reposTransport.CanRented,
		TransportType: reposTransport.Type,
		Model:         reposTransport.Model,
		Color:         reposTransport.Color,
		Identifier:    reposTransport.Identifier,
		Description:   sqlnt.ToString(&reposTransport.Description),
		Latitude:      reposTransport.Latitude,
		Longitude:     reposTransport.Longitude,
		MinutePrice:   sqlnt.ToF64(&reposTransport.MinutePrice),
		DayPrice:      sqlnt.ToF64(&reposTransport.MinutePrice),
		IsDeleted:     reposTransport.Deleted,
	}, nil
}

func (r *TransportPostgres) IsOwner(transportID, userID int32) (bool, error) {
	isOwner, err := r.queries.IsTransportOwner(context.Background(), queries.IsTransportOwnerParams{
		ID:    transportID,
		Owner: userID,
	})
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return isOwner, nil
}

func (r *TransportPostgres) Get(transportID int32) (*models.Transport, error) {
	result, err := r.queries.GetTransport(context.Background(), transportID)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Transport{
		ID:            result.ID,
		OwnerID:       result.Owner,
		CanBeRented:   result.CanRented,
		TransportType: result.Type,
		Model:         result.Model,
		Color:         result.Color,
		Identifier:    result.Identifier,
		Description:   sqlnt.ToString(&result.Description),
		Latitude:      result.Latitude,
		Longitude:     result.Longitude,
		MinutePrice:   sqlnt.ToF64(&result.MinutePrice),
		DayPrice:      sqlnt.ToF64(&result.DayPrice),
		IsDeleted:     result.Deleted,
	}, nil
}

func (r *TransportPostgres) IsExistByID(id int32) (bool, error) {
	isExist, err := r.queries.IsExistTransportByID(context.Background(), id)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *TransportPostgres) IsExistByIdentifier(identifier string) (bool, error) {
	isExist, err := r.queries.IsExistTransportByIdentifier(context.Background(), identifier)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *TransportPostgres) Create(transport *models.Transport) (*models.Transport, error) {
	result, err := r.queries.CreateTransport(context.Background(), queries.CreateTransportParams{
		Identifier:  transport.Identifier,
		Owner:       transport.OwnerID,
		Type:        transport.TransportType,
		CanRented:   transport.CanBeRented,
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
		ID:            result.ID,
		OwnerID:       result.Owner,
		CanBeRented:   result.CanRented,
		TransportType: result.Type,
		Model:         result.Model,
		Color:         result.Color,
		Identifier:    result.Identifier,
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
