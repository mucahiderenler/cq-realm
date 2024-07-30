package services

import (
	"mucahiderenler/conquerors-realm/internal/models"
	"mucahiderenler/conquerors-realm/internal/repository"
)

type MapService struct {
	Repo *repository.MapRepository
}

func NewMapService(repo *repository.MapRepository) *MapService {
	return &MapService{Repo: repo}
}

func (s *MapService) GetByID(mapID string) (*models.Map, error) {
	m, err := s.Repo.GetByID(mapID)

	if err != nil {
		return nil, err
	}

	return m, nil
}
