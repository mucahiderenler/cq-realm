package services

import (
	"encoding/json"
	"fmt"
	"os"
)

type Building struct {
	BuildingSpeed     map[int]int       `json:"buildingSpeed"`
	UpgradeTime       map[int]int       `json:"upgradeTime"`
	NeededPopulation  map[int]int       `json:"neededPopulation"`
	UpgradingCosts    map[int]Resources `json:"upgradingCosts"`
	PointByLevel      map[int]int       `json:"pointByLevel"`
	ProductionByLevel map[int]int       `json:"productionByLevel"`
}

type Unit struct {
}

type buildingName = string
type unitName = string

type GameConfig struct {
	Buildings map[buildingName]Building
	Units     map[unitName]Unit
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
