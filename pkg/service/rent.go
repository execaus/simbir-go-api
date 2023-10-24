package service

import (
	"simbir-go-api/pkg/repository"
)

type RentService struct {
	repo repository.Rent
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
