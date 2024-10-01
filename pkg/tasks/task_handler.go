package tasks

import (
	"mucahiderenler/conquerors-realm/internal/types"
	"mucahiderenler/conquerors-realm/internal/worker"

	"github.com/hibiken/asynq"
)

type TaskHandler struct {
	buildingUpgrade *worker.BuildingUpgrade
}

func NewTaskHandler(b *worker.BuildingUpgrade) *TaskHandler {
	return &TaskHandler{buildingUpgrade: b}
}

func (t *TaskHandler) GetTaskHandler() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.HandleFunc(types.TypeBuildingUpgrade, t.buildingUpgrade.HandleBuildingUpgradeTask)

	return mux
}
