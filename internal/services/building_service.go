package services

import (
	"context"
	"fmt"
	"mucahiderenler/conquerors-realm/internal/models"
	"mucahiderenler/conquerors-realm/internal/repository"
	"time"
)

// TODO:how about an env variable brotherman
const redisAddr = "172.17.0.4:6379"

type BuildingService struct {
	buildingRepo      *repository.BuildingRepository
	resourceService   *ResourceService
	gameConfigService *GameConfigService
}

type BuildingDetails struct {
	UpgradeTime      int              `json:"upgradeTime"`
	UpgradeCosts     models.Resources `json:"upgradeCosts"`
	CurrentLevel     int              `json:"currentLevel"`
	BuildingSpeed    int              `json:"buildingSpeed"`
	NeededPopulation int              `json:"neededPopulation"`
}

func NewBuildingService(resourceService *ResourceService,
	buildingRepo *repository.BuildingRepository,
	gameConfigService *GameConfigService,
) *BuildingService {
	return &BuildingService{resourceService: resourceService,
		buildingRepo:      buildingRepo,
		gameConfigService: gameConfigService,
	}
}

func (b *BuildingService) GetBuildingDetails(ctx context.Context, buildingId string, villageId string) (*BuildingDetails, error) {
	building, err := b.buildingRepo.GetVillageBuilding(ctx, buildingId, villageId)

	if err != nil {
		return nil, err
	}

	buildingConfig, ok := b.gameConfigService.GetBuildingConfig(building.Name)
	if !ok {
		return nil, fmt.Errorf("cannot find building config for: %s", building.Name)
	}

	curLevel := building.Level

	upgradeTime := buildingConfig.UpgradeTime[curLevel+1]
	upgradeCosts := buildingConfig.UpgradingCosts[curLevel+1]
	buildingSpeed := buildingConfig.BuildingSpeed[curLevel+1]
	neededPopulation := buildingConfig.NeededPopulation[curLevel+1]

	return &BuildingDetails{UpgradeTime: upgradeTime, UpgradeCosts: upgradeCosts, CurrentLevel: curLevel, BuildingSpeed: buildingSpeed, NeededPopulation: neededPopulation}, nil
}

func (b *BuildingService) UpgradeBuildingInit(ctx context.Context, buildingId string, villageId string) error {
	building, err := b.buildingRepo.GetVillageBuilding(ctx, buildingId, villageId)

	if err != nil {
		return err
	}

	buildingConfig, ok := b.gameConfigService.GetBuildingConfig(building.Name)

	if !ok {
		return fmt.Errorf("cannot find building config for: %s", building.Name)
	}

	currentResources, err := b.resourceService.GetVillageResources(ctx, villageId)

	if err != nil {
		return err
	}

	// assumes upgrade is next level, should we also request upgrade level? that also needs to be validated before starting upgrade
	var nextUpgradeLevel = building.Level + 1

	neededResources := buildingConfig.UpgradingCosts[nextUpgradeLevel]
	upgradeTime := buildingConfig.UpgradeTime[nextUpgradeLevel]

	isResourcesEnough := checkResources(neededResources, *currentResources)

	if !isResourcesEnough {
		return fmt.Errorf("resources are not enough for this upgrade %s", building.Name)
	}

	// start upgrading, decrease the resources from village
	currentResources.Clay -= neededResources.Clay
	currentResources.Iron -= neededResources.Iron
	currentResources.Wood -= neededResources.Wood

	b.buildingRepo.InsertResourcesBack(ctx, villageId, currentResources, time.Now())
	BuildingUpgradeTask(villageId, buildingId, upgradeTime)
	return nil

}

func (b *BuildingService) UpgradeBuilding(ctx context.Context, buildingId string, villageId string) error {
	building, err := b.buildingRepo.GetVillageBuilding(ctx, buildingId, villageId)

	if err != nil {
		return err
	}

	buildingConfig, ok := b.gameConfigService.GetBuildingConfig(building.Name)

	if !ok {
		return fmt.Errorf("cannot find building config for: %s", building.Name)
	}

	newBuildingLevel := building.Level + 1
	newProductionRate := buildingConfig.HourlyProductionByLevel[newBuildingLevel]

	err = b.buildingRepo.UpgradeBuilding(villageId, buildingId, newBuildingLevel, newProductionRate)

	if err != nil {
		return err
	}

	return nil
}

func checkResources(neededResources models.Resources, currentResources models.Resources) bool {
	if currentResources.Clay >= neededResources.Clay && currentResources.Iron >= neededResources.Iron && currentResources.Wood >= neededResources.Wood {
		return true
	}

	return false
}
