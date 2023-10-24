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

func (s *TransportService) GetFromRadius(point *models.Point, radius float64, transportType string) ([]models.Transport, error) {
	var err error
	var transports []models.Transport
	var reposTransports []queries.Transport

	if transportType == constants.TransportTypeAll {
		reposTransports, err = s.repo.GetFromRadiusAll(point, radius, transportType)
		if err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
	} else {
		reposTransports, err = s.repo.GetFromRadiusOnlyType(point, radius, transportType)
		if err != nil {
			exloggo.Error(err.Error())
			return nil, err
		}
	}

	for _, reposTransport := range reposTransports {
		transports = append(transports, models.Transport{
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
			OwnerID:       transport.Owner,
			CanBeRented:   transport.CanRanted,
			TransportType: transport.Type,
			Model:         transport.Model,
			Color:         transport.Color,
			Identifier:    transport.ID,
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

func (s *TransportService) IsRemoved(identifier string) (bool, error) {
	return s.repo.IsRemoved(identifier)
}

func (s *TransportService) Remove(identifier string) error {
	return s.repo.Remove(identifier)
}

func (s *TransportService) Update(identifier string, transport *models.Transport) (*models.Transport, error) {
	return s.repo.Update(identifier, transport)
}

func (s *TransportService) IsOwner(identifier, username string) (bool, error) {
	return s.repo.IsOwner(identifier, username)
}

func (s *TransportService) Get(identifier string) (*models.Transport, error) {
	return s.repo.Get(identifier)
}

func (s *TransportService) IsExist(identifier string) (bool, error) {
	return s.repo.IsExist(identifier)
}

func (s *TransportService) Create(transport *models.Transport) (*models.Transport, error) {
	return s.repo.Create(transport)
}

func NewTransportService(repo repository.TransportRepository) *TransportService {
	return &TransportService{repo: repo}
}
