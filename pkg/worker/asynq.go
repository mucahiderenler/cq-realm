package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const redisAddr = "172.17.0.4:6379"

func NewAsynqServer(lc fx.Lifecycle, taskHandler *TaskHandler, logger *zap.Logger) *asynq.Server {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{Concurrency: 10},
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				mux := taskHandler.GetTaskHandler()
				if err := server.Start(mux); err != nil {
					logger.Fatal("Error on starting worker,", zap.Error(err))
				}
			}()
			logger.Info("Asynq worker has been started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Shutdown()
			logger.Info("Asynq worker has been stopped")
			return nil
		},
	})

	return server
}
