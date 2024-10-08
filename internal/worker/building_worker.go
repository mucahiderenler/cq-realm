package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"mucahiderenler/conquerors-realm/internal/services"
	"mucahiderenler/conquerors-realm/internal/types"

	"github.com/hibiken/asynq"
)

type BuildingUpgrade struct {
	buildingService *services.BuildingService
}

func NewBuildingWorker(buildingService *services.BuildingService) *BuildingUpgrade {
	return &BuildingUpgrade{buildingService: buildingService}
}

func (b *BuildingUpgrade) HandleBuildingUpgradeTask(ctx context.Context, t *asynq.Task) error {
	var p types.BuildingUpgradePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		fmt.Println("Error happened while handling building task:", err)
		return err
	}

	err := b.buildingService.UpgradeBuilding(ctx, p.BuildingID, p.VillageID)

	if err != nil {
		fmt.Println("Something happened while upgrading building:", err)
		return err
	}

	fmt.Println("task finished successfully", p.BuildingID, p.VillageID)
	return nil
}
