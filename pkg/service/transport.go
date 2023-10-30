package service

import (
	"github.com/execaus/exloggo"
	"simbir-go-api/constants"
	"simbir-go-api/models"
	"simbir-go-api/pkg/repository"
	"simbir-go-api/pkg/repository/sqlnt"
	"simbir-go-api/queries"
)

type TransportService struct {
	repo repository.TransportRepository
}

const metersInMile = 1609

func (s *TransportService) IsAccessRent(id int32) (bool, error) {
	transport, err := s.repo.Get(id)
	if err != nil {
		exloggo.Error(err.Error())
		return false, err
	}

	return transport.CanBeRented, nil
}

func (s *TransportService) GetFromRadius(point *models.Point, radius float64, transportType string) ([]models.Transport, error) {
	var err error
	var transports []models.Transport
	var reposTransports []queries.Transport

	radiusForMile := radius * metersInMile

	if transportType == constants.TransportTypeAll {
		reposTransports, err = s.repo.GetFromRadiusAll(point, radiusForMile)
		if err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
	} else {
		reposTransports, err = s.repo.GetFromRadiusOnlyType(point, radiusForMile, transportType)
		if err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
	}

	for _, reposTransport := range reposTransports {
		transports = append(transports, models.Transport{
			ID:            reposTransport.ID,
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
			DayPrice:      sqlnt.ToF64(&reposTransport.DayPrice),
			IsDeleted:     reposTransport.Deleted,
		})
	}

	return transports, nil
}

func (s *TransportService) GetList(start, count int32, transportType string) ([]models.Transport, error) {
	var err error
	var reposTransports []queries.Transport
	transports := make([]models.Transport, 0)

	if transportType == constants.TransportTypeAll {
		reposTransports, err = s.repo.GetList(start, count)
		if err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
	} else {
		reposTransports, err = s.repo.GetListOnlyType(start, count, transportType)
		if err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
	}

	for _, transport := range reposTransports {
		transports = append(transports, models.Transport{
			ID:            transport.ID,
			OwnerID:       transport.Owner,
			CanBeRented:   transport.CanRented,
			TransportType: transport.Type,
			Model:         transport.Model,
			Color:         transport.Color,
			Identifier:    transport.Identifier,
			Description:   sqlnt.ToString(&transport.Description),
			Latitude:      transport.Latitude,
			Longitude:     transport.Longitude,
			MinutePrice:   sqlnt.ToF64(&transport.MinutePrice),
			DayPrice:      sqlnt.ToF64(&transport.DayPrice),
			IsDeleted:     transport.Deleted,
		})
	}

	return transports, nil
}

func (s *TransportService) IsRemoved(id int32) (bool, error) {
	return s.repo.IsRemoved(id)
}

func (s *TransportService) Remove(id int32) error {
	return s.repo.Remove(id)
}

func (s *TransportService) Update(transport *models.Transport) (*models.Transport, error) {
	return s.repo.Update(transport)
}

func (s *TransportService) IsOwner(id, userID int32) (bool, error) {
	return s.repo.IsOwner(id, userID)
}

func (s *TransportService) Get(id int32) (*models.Transport, error) {
	return s.repo.Get(id)
}

func (s *TransportService) IsExistByID(id int32) (bool, error) {
	return s.repo.IsExistByID(id)
}

func (s *TransportService) IsExistByIdentifier(identifier string) (bool, error) {
	return s.repo.IsExistByIdentifier(identifier)
}

func (s *TransportService) Create(transport *models.Transport) (*models.Transport, error) {
	return s.repo.Create(transport)
}

func NewTransportService(repo repository.TransportRepository) *TransportService {
	return &TransportService{repo: repo}
}
