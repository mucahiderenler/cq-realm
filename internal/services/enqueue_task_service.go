package services

import (
	"encoding/json"
	"fmt"
	"mucahiderenler/conquerors-realm/internal/types"
	"time"

	"github.com/hibiken/asynq"
)

func BuildingUpgradeTask(villageID string, buildingID string, upgradeTime int) error {
	payload, err := json.Marshal(types.BuildingUpgradePayload{VillageID: villageID, BuildingID: buildingID})
	if err != nil {
		fmt.Println("Building upgrade task json marshal failed", err)
		return err
	}

	fmt.Println("Building upgrading task initializing: ", string(payload))

	task := asynq.NewTask(types.TypeBuildingUpgrade, payload)

	enqueueTask(task, upgradeTime)
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
