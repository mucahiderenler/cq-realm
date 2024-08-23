package services

import (
	"context"
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
	_, err := b.buildingRepo.GetBuildingById(ctx, buildingId)

	fmt.Println(b.gameConfigService.config.Buildings["barracks"].UpgradeTime)
	if err != nil {
		return err
	}

	return nil

	// give building id and level to game config service, it will return the conditons for upgrade
	// conditons, err := b.gameConfigService.getConditions(ctx, building)

	// if conditons met (for now just resources, in the future maybe other buildings will effect the upgrade of a building)

	// b.resourceService.GetVillageResources()
}
