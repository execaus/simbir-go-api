package service

import (
	"simbir-go-api/models"
	"simbir-go-api/pkg/repository"
)

type TransportService struct {
	repo repository.TransportRepository
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
