package repository

import (
	"context"
	"mucahiderenler/conquerors-realm/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var resourceBuildingTypes = []interface{}{2, 3, 5} // 2 - Iron mine, 3 - Lumberjack, 5 - Clay pit

type BuildingRepository struct {
	DB *sqlx.DB
}

func NewBuildingRepository(DB *sqlx.DB) *BuildingRepository {
	return &BuildingRepository{DB: DB}
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

func (r *BuildingRepository) InsertResourcesBack(ctx context.Context, buildingId int, currentResouce float64, now time.Time) {
	building, _ := models.FindBuilding(ctx, r.DB, buildingId)

	building.LastInteraction = null.TimeFrom(now)
	building.LastResource = null.Float64From(currentResouce)

	building.Update(ctx, r.DB, boil.Infer())
}
