package services

import (
	"mucahiderenler/conquerors-realm/internal/models"
)

type MapService struct {
	villageService *VillageService
}

func NewMapService(villageService *VillageService) *MapService {
	return &MapService{villageService: villageService}
}

func (s *MapService) GetMap() (*models.Map, error) {
	villages, err := s.villageService.GetAllVillages()
	if err != nil {
		return nil, err
	}

	m := &models.Map{}
	m.Villages = villages

	return m, nil
}
