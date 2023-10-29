package service

import (
	"github.com/execaus/exloggo"
	"simbir-go-api/constants"
	"simbir-go-api/models"
	"simbir-go-api/pkg/repository"
	"simbir-go-api/pkg/repository/sqlnt"
	"time"
)

type RentService struct {
	repo repository.Rent
}

const dayHours = float64(time.Hour * 24)

func (s *RentService) Remove(id int32) error {
	return s.repo.Remove(id)
}

func (s *RentService) Update(rent *models.Rent) (*models.Rent, error) {
	updatedRent, err := s.repo.Update(rent)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	endTime := sqlnt.ToTime(&updatedRent.TimeEnd)

	return &models.Rent{
		ID:         updatedRent.ID,
		Account:    updatedRent.Account,
		Transport:  updatedRent.Transport,
		TimeStart:  updatedRent.TimeStart,
		TimeEnd:    endTime,
		PriceUnit:  updatedRent.PriceUnit,
		PriceType:  updatedRent.PriceType,
		FinalPrice: calculateFinalPrice(updatedRent.PriceType, updatedRent.PriceUnit, updatedRent.TimeStart, endTime),
		IsDeleted:  updatedRent.Deleted,
	}, nil
}

func (s *RentService) End(id int32) (*models.Rent, error) {
	endTime := time.Now()

	if err := s.repo.End(id, &endTime); err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	rent, err := s.repo.Get(id)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &models.Rent{
		ID:         rent.ID,
		Account:    rent.Account,
		Transport:  rent.Transport,
		TimeStart:  rent.TimeStart,
		TimeEnd:    sqlnt.ToTime(&rent.TimeEnd),
		PriceUnit:  rent.PriceUnit,
		PriceType:  rent.PriceType,
		FinalPrice: calculateFinalPrice(rent.PriceType, rent.PriceUnit, rent.TimeStart, &endTime),
		IsDeleted:  rent.Deleted,
	}, nil
}

func (s *RentService) Create(rent *models.Rent) (*models.Rent, error) {
	createdRent, err := s.repo.Create(rent)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	timeEnd := sqlnt.ToTime(&createdRent.TimeEnd)

	return &models.Rent{
		ID:         rent.ID,
		Account:    rent.Account,
		Transport:  rent.Transport,
		TimeStart:  rent.TimeStart,
		TimeEnd:    timeEnd,
		PriceUnit:  rent.PriceUnit,
		PriceType:  rent.PriceType,
		IsDeleted:  rent.IsDeleted,
		FinalPrice: calculateFinalPrice(createdRent.PriceType, createdRent.PriceUnit, createdRent.TimeStart, timeEnd),
	}, nil
}

func (s *RentService) TransportIsRented(transportID int32) (bool, error) {
	isExistRent, err := s.repo.IsExistCurrentRented(transportID)
	if err != nil {
		exloggo.Error(err.Error())
		return false, nil
	}

	return isExistRent, nil
}

func (s *RentService) GetListFromTransportID(transportID, start, count int32) ([]models.Rent, error) {
	var rents []models.Rent

	reposRents, err := s.repo.GetListFromTransportID(transportID, start, count)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	for _, reposResult := range reposRents {
		endTime := sqlnt.ToTime(&reposResult.TimeEnd)
		rents = append(rents, models.Rent{
			ID:         reposResult.ID,
			Account:    reposResult.Account,
			Transport:  reposResult.Transport,
			TimeStart:  reposResult.TimeStart,
			TimeEnd:    endTime,
			PriceUnit:  reposResult.PriceUnit,
			PriceType:  reposResult.PriceType,
			FinalPrice: calculateFinalPrice(reposResult.PriceType, reposResult.PriceUnit, reposResult.TimeStart, endTime),
			IsDeleted:  reposResult.Deleted,
		})
	}

	return rents, nil
}

func (s *RentService) GetListFromUserID(userID, start, count int32) ([]models.Rent, error) {
	var rents []models.Rent

	reposRents, err := s.repo.GetListFromUserID(userID, start, count)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}

	for _, reposResult := range reposRents {
		endTime := sqlnt.ToTime(&reposResult.TimeEnd)
		rents = append(rents, models.Rent{
			ID:         reposResult.ID,
			Account:    reposResult.Account,
			Transport:  reposResult.Transport,
			TimeStart:  reposResult.TimeStart,
			TimeEnd:    endTime,
			PriceUnit:  reposResult.PriceUnit,
			PriceType:  reposResult.PriceType,
			FinalPrice: calculateFinalPrice(reposResult.PriceType, reposResult.PriceUnit, reposResult.TimeStart, endTime),
			IsDeleted:  reposResult.Deleted,
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

	endTime := sqlnt.ToTime(&reposResult.TimeEnd)

	return &models.Rent{
		ID:         reposResult.ID,
		Account:    reposResult.Account,
		Transport:  reposResult.Transport,
		TimeStart:  reposResult.TimeStart,
		TimeEnd:    endTime,
		PriceUnit:  reposResult.PriceUnit,
		PriceType:  reposResult.PriceType,
		FinalPrice: calculateFinalPrice(reposResult.PriceType, reposResult.PriceUnit, reposResult.TimeStart, endTime),
		IsDeleted:  reposResult.Deleted,
	}, nil
}

func (s *RentService) IsRenter(id int32, userID int32) (bool, error) {
	return s.repo.IsRenter(id, userID)
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

func calculateFinalPrice(rentType string, rentUnit float64, startTime time.Time, endTime *time.Time) *float64 {
	if endTime == nil {
		return nil
	}

	var duration time.Duration
	var result float64
	if rentType == constants.RentTypeMinutes {
		duration = (*endTime).Sub(startTime)
		minutes := duration.Minutes()
		result = minutes * rentUnit
		return &result
	} else if rentType == constants.RentTypeDays {
		duration = (*endTime).Sub(startTime)
		days := duration.Hours()
		result = days / dayHours * rentUnit
		return &result
	}

	exloggo.Error("invalid rent type")
	return nil
}
