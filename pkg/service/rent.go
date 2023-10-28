package service

import (
	"github.com/execaus/exloggo"
	"simbir-go-api/models"
	"simbir-go-api/pkg/repository"
	"simbir-go-api/pkg/repository/sqlnt"
)

type RentService struct {
	repo repository.Rent
}

func (s *RentService) GetListFromTransport(transportID string) ([]models.Rent, error) {
	var rents []models.Rent

	reposRents, err := s.repo.GetListFromTransport(transportID)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	for _, reposResult := range reposRents {
		rents = append(rents, models.Rent{
			ID: reposResult.ID,
			Account: models.Account{
				Username:  reposResult.Username,
				Password:  "",
				Balance:   reposResult.Balance,
				Roles:     []string{},
				IsDeleted: reposResult.Deleted_2,
			},
			Transport: models.Transport{
				OwnerID:       reposResult.Owner,
				CanBeRented:   reposResult.CanRanted,
				TransportType: reposResult.Type,
				Model:         reposResult.Model,
				Color:         reposResult.Color,
				Identifier:    reposResult.ID_2,
				Description:   sqlnt.ToString(&reposResult.Description),
				Latitude:      reposResult.Latitude,
				Longitude:     reposResult.Longitude,
				MinutePrice:   sqlnt.ToF64(&reposResult.MinutePrice),
				DayPrice:      sqlnt.ToF64(&reposResult.DayPrice),
				IsDeleted:     reposResult.Deleted_3,
			},
			TimeStart: reposResult.TimeStart,
			TimeEnd:   sqlnt.ToTime(&reposResult.TimeEnd),
			PriceUnit: reposResult.PriceUnit,
			PriceType: reposResult.PriceType,
			IsDeleted: reposResult.Deleted,
		})
	}

	return rents, nil
}

func (s *RentService) GetListFromUsername(username string) ([]models.Rent, error) {
	var rents []models.Rent

	reposRents, err := s.repo.GetListFromUsername(username)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	for _, reposResult := range reposRents {
		rents = append(rents, models.Rent{
			ID: reposResult.ID,
			Account: models.Account{
				Username:  reposResult.Username,
				Password:  "",
				Balance:   reposResult.Balance,
				Roles:     []string{},
				IsDeleted: reposResult.Deleted_2,
			},
			Transport: models.Transport{
				OwnerID:       reposResult.Owner,
				CanBeRented:   reposResult.CanRanted,
				TransportType: reposResult.Type,
				Model:         reposResult.Model,
				Color:         reposResult.Color,
				Identifier:    reposResult.ID_2,
				Description:   sqlnt.ToString(&reposResult.Description),
				Latitude:      reposResult.Latitude,
				Longitude:     reposResult.Longitude,
				MinutePrice:   sqlnt.ToF64(&reposResult.MinutePrice),
				DayPrice:      sqlnt.ToF64(&reposResult.DayPrice),
				IsDeleted:     reposResult.Deleted_3,
			},
			TimeStart: reposResult.TimeStart,
			TimeEnd:   sqlnt.ToTime(&reposResult.TimeEnd),
			PriceUnit: reposResult.PriceUnit,
			PriceType: reposResult.PriceType,
			IsDeleted: reposResult.Deleted,
		})
	}

	return rents, nil
}

func (s *RentService) Get(id int32) (*models.Rent, error) {
	reposResult, err := s.repo.Get(id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Rent{
		ID: reposResult.ID,
		Account: models.Account{
			Username:  reposResult.Username,
			Password:  "",
			Balance:   reposResult.Balance,
			Roles:     []string{},
			IsDeleted: reposResult.Deleted_2,
		},
		Transport: models.Transport{
			OwnerID:       reposResult.Owner,
			CanBeRented:   reposResult.CanRanted,
			TransportType: reposResult.Type,
			Model:         reposResult.Model,
			Color:         reposResult.Color,
			Identifier:    reposResult.ID_2,
			Description:   sqlnt.ToString(&reposResult.Description),
			Latitude:      reposResult.Latitude,
			Longitude:     reposResult.Longitude,
			MinutePrice:   sqlnt.ToF64(&reposResult.MinutePrice),
			DayPrice:      sqlnt.ToF64(&reposResult.DayPrice),
			IsDeleted:     reposResult.Deleted_3,
		},
		TimeStart: reposResult.TimeStart,
		TimeEnd:   sqlnt.ToTime(&reposResult.TimeEnd),
		PriceUnit: reposResult.PriceUnit,
		PriceType: reposResult.PriceType,
		IsDeleted: reposResult.Deleted,
	}, nil
}

func (s *RentService) IsRenter(id int32, username string) (bool, error) {
	return s.repo.IsRenter(id, username)
}

func (s *RentService) IsExist(id int32) (bool, error) {
	return s.repo.IsExist(id)
}

func (s *RentService) IsRemoved(id int32) (bool, error) {
	return s.repo.IsRemoved(id)
}

func NewRentService(repo repository.Rent) *RentService {
	return &RentService{repo: repo}
}
