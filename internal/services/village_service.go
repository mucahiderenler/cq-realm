package services

import (
	"context"
	models "mucahiderenler/conquerors-realm/internal/models"
	"mucahiderenler/conquerors-realm/internal/repository"

	"go.uber.org/zap"
)

type VillageService struct {
	Repo   *repository.VillageRepository
	Logger *zap.Logger
}

type GetVillageByIDResult struct {
	Village   *models.Village
	Buildings []*models.Building
}

func NewVillageService(repo *repository.VillageRepository, logger *zap.Logger) *VillageService {
	return &VillageService{Repo: repo, Logger: logger}
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
	buildings := village.R.Buildings
	result := &GetVillageByIDResult{Village: village, Buildings: buildings}
	if err != nil {
		return nil, err
	}
	// Perform any additional business logic if needed
	return result, nil
}
