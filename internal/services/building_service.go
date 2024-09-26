package services

import (
	"context"
	"errors"
	"fmt"
	models "mucahiderenler/conquerors-realm/internal/models"
	"mucahiderenler/conquerors-realm/internal/repository"
	"mucahiderenler/conquerors-realm/internal/tasks"
	"time"
)

type BuildingService struct {
	buildingRepo      *repository.BuildingRepository
	resourceService   *ResourceService
	gameConfigService *GameConfigService
	taskHandler       *tasks.TaskHandler
}

func NewBuildingService(resourceService *ResourceService,
	buildingRepo *repository.BuildingRepository,
	gameConfigService *GameConfigService,
	taskHandler *tasks.TaskHandler) *BuildingService {
	return &BuildingService{resourceService: resourceService,
		buildingRepo:      buildingRepo,
		gameConfigService: gameConfigService,
		taskHandler:       taskHandler,
	}
}

func (b *BuildingService) UpgradeBuildingInit(ctx context.Context, buildingId string, villageId string) error {
	building, err := b.buildingRepo.GetBuildingById(ctx, buildingId)

	if err != nil {
		return err
	}

	buildingConfig, ok := b.gameConfigService.GetBuildingConfig(building.Name)

	if !ok {
		return errors.New(fmt.Sprintf("Cannot find the building config for: %s ", building.Name))
	}

	currentResources, err := b.resourceService.GetVillageResources(ctx, villageId)

	if err != nil {
		return err
	}

	// assumes upgrade is next level, should we also request upgrade level? that also needs to be validated before starting upgrade
	var nextUpgradeLevel = building.Level + 1

	neededResources := buildingConfig.UpgradingCosts[nextUpgradeLevel]

	isResourcesEnough := checkResources(neededResources, *currentResources)

	if !isResourcesEnough {
		return errors.New(fmt.Sprintf("Resources are not enough for this upgrade", building.Name))
	}

	// start upgrading, decrease the resources from village
	currentResources.Clay -= neededResources.Clay
	currentResources.Iron -= neededResources.Iron
	currentResources.Wood -= neededResources.Wood
	b.buildingRepo.InsertResourcesBack(ctx, currentResources, time.Now())
	b.taskHandler.BuildingUpgradeTask("1", "2", "10")
	return nil

}

func checkResources(neededResources models.Resources, currentResources models.Resources) bool {
	if currentResources.Clay >= neededResources.Clay && currentResources.Iron >= neededResources.Iron && currentResources.Wood >= neededResources.Wood {
		return true
	}

	return false
}
