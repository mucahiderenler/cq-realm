package repository

import (
	"context"
	"mucahiderenler/conquerors-realm/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var resourceBuildingTypes = []interface{}{2, 3, 5} // 2 - Iron mine, 3 - Lumberjack, 5 - Clay pit

type BuildingRepository struct {
	DB *sqlx.DB
}

func NewBuildingRepository(DB *sqlx.DB) *BuildingRepository {
	return &BuildingRepository{DB: DB}
}

func (r *BuildingRepository) GetVillageBuilding(ctx context.Context, buildingId string, villageId string) (*models.Building, error) {
	building, err := models.Buildings(Where("id = ?", buildingId), (Where("village_id = ?", villageId))).One(ctx, r.DB)

	if err != nil {
		return nil, err
	}

	return building, nil
}

func (r *BuildingRepository) GetResourceBuildingsForVillage(ctx context.Context, villageId string) ([]*models.Building, error) {
	var resourceBuildings []*models.Building
	resourceBuildings, err := models.Buildings(Where("village_id = ?", villageId), WhereIn("building_type IN ?", resourceBuildingTypes...)).All(ctx, r.DB)

	if err != nil {
		return nil, err
	}

	return resourceBuildings, nil
}

func (r *BuildingRepository) GetStorageBuildingForVillage(ctx context.Context, villageId string) (*models.Building, error) {
	var storageBuilding *models.Building

	storageBuilding, err := models.Buildings(Where("village_id = ? AND building_type = ?", villageId, 6)).One(ctx, r.DB)

	if err != nil {
		return nil, err
	}

	return storageBuilding, nil
}

func (r *BuildingRepository) InsertResourcesBack(ctx context.Context, villageId string, resources *models.Resources, now time.Time) error {
	query := `update buildings set 
	last_resource = CASE
		WHEN building_type = 2 THEN $1
		WHEN building_type = 3 THEN $2 
		WHEN building_type = 5 THEN $3 
		ELSE last_resource
	END,

	last_interaction = $4 


	where village_id = $5 and building_type in (2,3,5)`

	_, err := r.DB.Exec(query, resources.Iron, resources.Wood, resources.Clay, now, villageId)

	if err != nil {
		return err
	}

	return nil
}

func (r *BuildingRepository) UpgradeBuilding(villageId string, buildingId string, newLevel int, newProductionRate int) error {
	query := `update buildings set production_rate = $1, level = $2 where village_id = $3 and building_type = $4`

	_, err := r.DB.Exec(query, newProductionRate, newLevel, villageId, buildingId)

	if err != nil {
		return err
	}

	return nil
}
