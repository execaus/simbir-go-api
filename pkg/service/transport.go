package service

import (
	"simbir-go-api/models"
	"simbir-go-api/pkg/repository"
)

type TransportService struct {
	repo repository.TransportRepository
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
