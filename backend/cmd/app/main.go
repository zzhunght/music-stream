package main

import (
	"music-app-backend/internal/app/api"
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/utils"
	"music-app-backend/worker"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func main() {
	env, err := utils.LoadEnv(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	redisOpt := asynq.RedisClientOpt{
		Addr: env.RedisUrl,
	}
	taskClient := worker.NewDeliveryTaskClient(redisOpt)
	go StartRedisWorker(redisOpt)
	store := config.InitDB(env.DatabaseDestination)
	server := api.NewServer(store, env, taskClient)
	// defer config.CloseDB()
	server.Run(":8080")
}

func StartRedisWorker(redisOpts asynq.RedisClientOpt) {
	log.Info().Msg("Start Task processor")
	client := worker.NewProcessorTaskClient(redisOpts)
	err := client.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start worker")
	}
}
