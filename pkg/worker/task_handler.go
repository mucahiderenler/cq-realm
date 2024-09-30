package worker

import (
	"mucahiderenler/conquerors-realm/internal/tasks"
	"mucahiderenler/conquerors-realm/internal/types"

	"github.com/hibiken/asynq"
)

type TaskHandler struct {
	buildingUpgradeHandler *tasks.BuildingUpgradeHandler
}

func NewTaskHandler(b *tasks.BuildingUpgradeHandler) *TaskHandler {
	return &TaskHandler{buildingUpgradeHandler: b}
}

func (t *TaskHandler) GetTaskHandler() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.HandleFunc(types.TypeBuildingUpgrade, t.buildingUpgradeHandler.HandleBuildingUpgradeTask)

	return mux
}
