package services

import (
	"context"
	"mucahiderenler/conquerors-realm/internal/models"
)

type MapService struct {
	villageService *VillageService
}

func NewMapService(villageService *VillageService) *MapService {
	return &MapService{villageService: villageService}
}

func (s *MapService) GetMap(ctx context.Context) (*models.Map, error) {
	villages, err := s.villageService.GetAllVillages(ctx)
	if err != nil {
		return nil, err
	}

	m := &models.Map{}

	m.Villages = villages
	return m, nil
}
