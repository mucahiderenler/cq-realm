package services

import (
	"context"
	models "mucahiderenler/conquerors-realm/internal/models"
	"mucahiderenler/conquerors-realm/internal/repository"
	"time"

	"go.uber.org/zap"
)

type ResourceService struct {
	BuildingRepo *repository.BuildingRepository
	Logger       *zap.Logger
}

func NewResourceService(buildingRepo *repository.BuildingRepository, logger *zap.Logger) *ResourceService {
	return &ResourceService{BuildingRepo: buildingRepo, Logger: logger}
}

func (s *ResourceService) GetVillageResources(ctx context.Context, villageId string) (*models.Resources, error) {
	resourceBuildings, err := s.BuildingRepo.GetResourceBuildingsForVillage(ctx, villageId)

	if err != nil {
		return nil, err
	}

	storageBuilding, err := s.BuildingRepo.GetStorageBuildingForVillage(ctx, villageId)

	if err != nil {
		return nil, err
	}

	var resources *models.Resources = &models.Resources{Wood: 0, Clay: 0, Iron: 0}
	now := time.Now()
	for _, resourceBuilding := range resourceBuildings {
		minutesPast := now.Sub(resourceBuilding.LastInteraction.Time).Minutes()
		currentResource := resourceBuilding.LastResource.Float64 + (minutesPast * resourceBuilding.ProductionRate.Float64)

		// if calculated resource is bigger than our storage capacity, set the resource to capacity
		if currentResource > storageBuilding.ProductionRate.Float64 {
			currentResource = storageBuilding.ProductionRate.Float64
		}

		if resourceBuilding.BuildingType == 2 {
			resources.Iron = int(currentResource)
		} else if resourceBuilding.BuildingType == 3 {
			resources.Wood = int(currentResource)
		} else {
			resources.Clay = int(currentResource)
		}
	}

	// insert resources and last interaction back
	s.BuildingRepo.InsertResourcesBack(ctx, villageId, resources, now)
	return resources, nil
}
