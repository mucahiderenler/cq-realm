package services

import (
	"encoding/json"
	"fmt"
	models "mucahiderenler/conquerors-realm/internal/models"
	"os"
)

type BuildingConfig struct {
	BuildingSpeed           map[int]int              `json:"buildingSpeed"`
	UpgradeTime             map[int]int              `json:"upgradeTime"`
	NeededPopulation        map[int]int              `json:"neededPopulation"`
	UpgradingCosts          map[int]models.Resources `json:"upgradingCosts"`
	PointByLevel            map[int]int              `json:"pointByLevel"`
	HourlyProductionByLevel map[int]int              `json:"hourlyProductionByLevel"`
}

type Unit struct {
}

type buildingName = string
type unitName = string

type GameConfig struct {
	BuildingConfigs map[buildingName]BuildingConfig `json:"buildingConfigs"`
	Units           map[unitName]Unit               `json:"units"`
}

type GameConfigService struct {
	config GameConfig
}

func NewGameConfigService() *GameConfigService {
	service := &GameConfigService{}
	err := service.loadConfig("gameConfig.json")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return service
}

func (s *GameConfigService) loadConfig(filepath string) error {
	data, err := os.ReadFile(filepath)

	if err != nil {
		return err
	}

	var config GameConfig

	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	s.config = config
	return nil
}

func (s *GameConfigService) GetBuildingConfig(buildingName string) (*BuildingConfig, bool) {
	buildingConfig, ok := s.config.BuildingConfigs[buildingName]

	if ok {
		return &buildingConfig, ok
	}

	return nil, ok
}
