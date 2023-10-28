package service

import (
	"github.com/execaus/exloggo"
	"simbir-go-api/models"
	"simbir-go-api/pkg/repository"
	"simbir-go-api/pkg/repository/sqlnt"
	"time"
)

type RentService struct {
	repo repository.Rent
}

func (s *RentService) End(id int32) (*models.Rent, error) {
	timeEnd := time.Now()

	if err := s.repo.End(id, &timeEnd); err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	rent, err := s.repo.Get(id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Rent{
		ID: rent.ID,
		Account: models.Account{
			Username:  rent.Username,
			Password:  "",
			Balance:   rent.Balance,
			Roles:     []string{},
			IsDeleted: rent.Deleted_2,
		},
		Transport: models.Transport{
			OwnerID:       rent.Owner,
			CanBeRented:   rent.CanRanted,
			TransportType: rent.Type,
			Model:         rent.Model,
			Color:         rent.Color,
			Identifier:    rent.ID_2,
			Description:   sqlnt.ToString(&rent.Description),
			Latitude:      rent.Latitude,
			Longitude:     rent.Longitude,
			MinutePrice:   sqlnt.ToF64(&rent.MinutePrice),
			DayPrice:      sqlnt.ToF64(&rent.DayPrice),
			IsDeleted:     rent.Deleted_3,
		},
		TimeStart: rent.TimeStart,
		TimeEnd:   sqlnt.ToTime(&rent.TimeEnd),
		PriceUnit: rent.PriceUnit,
		PriceType: rent.PriceType,
		IsDeleted: rent.Deleted,
	}, nil
}

func (s *RentService) Create(username, transportID string, timeStart time.Time, timeEnd *time.Time, priceUnit float64, rentType string) (*models.Rent, error) {
	createdRent, err := s.repo.Create(username, transportID, timeStart, timeEnd, priceUnit, rentType)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	rent, err := s.repo.Get(createdRent.ID)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Rent{
		ID: rent.ID,
		Account: models.Account{
			Username:  rent.Username,
			Password:  "",
			Balance:   rent.Balance,
			Roles:     []string{},
			IsDeleted: rent.Deleted_2,
		},
		Transport: models.Transport{
			OwnerID:       rent.Owner,
			CanBeRented:   rent.CanRanted,
			TransportType: rent.Type,
			Model:         rent.Model,
			Color:         rent.Color,
			Identifier:    rent.ID_2,
			Description:   sqlnt.ToString(&rent.Description),
			Latitude:      rent.Latitude,
			Longitude:     rent.Longitude,
			MinutePrice:   sqlnt.ToF64(&rent.MinutePrice),
			DayPrice:      sqlnt.ToF64(&rent.DayPrice),
			IsDeleted:     rent.Deleted_3,
		},
		TimeStart: rent.TimeStart,
		TimeEnd:   sqlnt.ToTime(&rent.TimeEnd),
		PriceUnit: rent.PriceUnit,
		PriceType: rent.PriceType,
		IsDeleted: rent.Deleted,
	}, nil
}

func (s *RentService) TransportIsRented(transportID string) (bool, error) {
	isExistRent, err := s.repo.IsExistCurrentRented(transportID)
	if err != nil {
		exloggo.Error(err.Error())
		return false, nil
	}

	return isExistRent, nil
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
