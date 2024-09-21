package services

import (
	"context"
	"errors"
	"fmt"
	"mucahiderenler/conquerors-realm/internal/repository"
)

type BuildingService struct {
	buildingRepo      *repository.BuildingRepository
	resourceService   *ResourceService
	gameConfigService *GameConfigService
}

func NewBuildingService(resourceService *ResourceService, buildingRepo *repository.BuildingRepository, gameConfigService *GameConfigService) *BuildingService {
	return &BuildingService{resourceService: resourceService, buildingRepo: buildingRepo, gameConfigService: gameConfigService}
}

func (b *BuildingService) UpgradeBuildingInit(ctx context.Context, buildingId string, villageId string) error {
	building, err := b.buildingRepo.GetBuildingById(ctx, buildingId)

	if err != nil {
		return err
	}

	buildingConfig, ok := b.gameConfigService.GetBuildingConfig(building.Name)

	if !ok {
		return errors.New(fmt.Sprintf("Cannot find the building config for: ", building.Name))
	}

	return nil

	// give building id and level to game config service, it will return the conditons for upgrade
	// conditons, err := b.gameConfigService.getConditions(ctx, building)

	// if conditons met (for now just resources, in the future maybe other buildings will effect the upgrade of a building)

	// b.resourceService.GetVillageResources()
}
