package services

import (
	"context"
	"mucahiderenler/conquerors-realm/internal/repository"
	"time"

	"go.uber.org/zap"
)

type ResourceService struct {
	BuildingRepo *repository.BuildingRepository
	Logger       *zap.Logger
}

type Resources struct {
	Wood int `json:"wood"`
	Clay int `json:"clay"`
	Iron int `json:"iron"`
}

func NewResourceService(buildingRepo *repository.BuildingRepository, logger *zap.Logger) *ResourceService {
	return &ResourceService{BuildingRepo: buildingRepo, Logger: logger}
}

func (s *ResourceService) GetVillageResources(ctx context.Context, villageId string) (*Resources, error) {
	resourceBuildings, err := s.BuildingRepo.GetResourceBuildingsForVillage(ctx, villageId)

	if err != nil {
		return nil, err
	}

	storageBuilding, err := s.BuildingRepo.GetStorageBuildingForVillage(ctx, villageId)

	if err != nil {
		return nil, err
	}

	var resources *Resources = &Resources{Wood: 0, Clay: 0, Iron: 0}
	now := time.Now()
	for i := 0; i < len(resourceBuildings); i++ {
		resourceBuilding := resourceBuildings[i]
		minutesPast := now.Sub(resourceBuilding.LastInteraction.Time).Minutes()
		currentResource := resourceBuilding.LastResource.Float64 + (minutesPast * resourceBuilding.ProductionRate.Float64)

		if currentResource > storageBuilding.ProductionRate.Float64 {
			currentResource = storageBuilding.ProductionRate.Float64
		}

		// insert resources and last interaction back
		s.BuildingRepo.InsertResourcesBack(ctx, resourceBuilding.ID, currentResource, now)

		if resourceBuilding.BuildingType == 2 {
			resources.Iron = int(currentResource)
		} else if resourceBuilding.BuildingType == 3 {
			resources.Wood = int(currentResource)
		} else {
			resources.Clay = int(currentResource)
		}
	}

	return resources, nil
}
