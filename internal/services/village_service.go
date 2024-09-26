package services

import (
	"context"
	models "mucahiderenler/conquerors-realm/internal/models"
	"mucahiderenler/conquerors-realm/internal/repository"

	"go.uber.org/zap"
)

type VillageService struct {
	Repo            *repository.VillageRepository
	ResourceService *ResourceService
	Logger          *zap.Logger
}

type GetVillageByIDResult struct {
	Village   *models.Village    `json:"village"`
	Buildings []*models.Building `json:"buildings"`
	Resources *models.Resources  `json:"resources"`
}

func NewVillageService(repo *repository.VillageRepository, resouceService *ResourceService, logger *zap.Logger) *VillageService {
	return &VillageService{Repo: repo, ResourceService: resouceService, Logger: logger}
}

func (s *VillageService) GetAllVillages(ctx context.Context) ([]*models.Village, error) {
	villages, err := s.Repo.GetAllVillages(ctx)

	if err != nil {
		return nil, err
	}

	return villages, nil
}

// Get a village by its ID
func (s *VillageService) GetVillageByID(ctx context.Context, id string) (*GetVillageByIDResult, error) {
	village, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	buildings := village.R.Buildings
	resources, err := s.ResourceService.GetVillageResources(ctx, id)
	result := &GetVillageByIDResult{Village: village, Buildings: buildings, Resources: resources}
	if err != nil {
		return nil, err
	}
	// Perform any additional business logic if needed
	return result, nil
}
