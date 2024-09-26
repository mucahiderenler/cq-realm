package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

type TaskHandler struct{}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

const redisAddr = "172.17.0.2:6379"

const (
	TypeBuildingUpgrade = "building:upgrade"
	TyepUnitRecruit     = "unit:recruit"
)

type BuildingUpgradePayload struct {
	UserID     string
	VillageID  string
	BuildingID string
}

func (t *TaskHandler) GetTaskHandler() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeBuildingUpgrade, HandleBuildingUpgradeTask)

	return mux
}

func (t *TaskHandler) BuildingUpgradeTask(userID string, villageID string, buildingID string) error {
	payload, err := json.Marshal(BuildingUpgradePayload{UserID: userID, VillageID: villageID, BuildingID: buildingID})

	if err != nil {
		fmt.Println("Building upgrade task json marshal failed", err)
		return err
	}

	fmt.Println("Building upgrading task initializing: ", payload)

	task := asynq.NewTask(TypeBuildingUpgrade, payload)

	enqueueTask(task, 10)
	return nil
}

func enqueueTask(task *asynq.Task, seconds int) (*asynq.TaskInfo, error) {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	info, err := client.Enqueue(task, asynq.ProcessIn(time.Second*time.Duration(seconds)))

	if err != nil {
		fmt.Println("Error happened while enqueing the task", err)
		return nil, err
	}

	return info, nil
}

func HandleBuildingUpgradeTask(ctx context.Context, t *asynq.Task) error {
	var p BuildingUpgradePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		fmt.Println("Error happened while handling building task:", err)
		return err
	}

	fmt.Printf("userID: %s, VillageID: %s, BuildingID: %s, do the upgrade.", p.UserID, p.VillageID, p.VillageID)
	return nil
}
