package main

import (
	"music-app-backend/internal/app/api"
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/utils"
	"music-app-backend/message"
	"music-app-backend/sqlc"
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
	messageQueue, err := message.NewRabbitMQ(env.RabbitMQUrl)
	if err != nil {
		log.Fatal().Err(err).Msg(("can connect to message queue"))
	}
	err = messageQueue.DeclareExchange()
	if err != nil {
		log.Fatal().Err(err).Msg(("can declare exchange"))
	}
	defer messageQueue.CloseRabbitMQ()
	taskClient := worker.NewDeliveryTaskClient(redisOpt)
	mailsender := utils.NewMailSender(env)
	store := config.InitDB(env.DatabaseDestination)
	go StartRedisWorker(redisOpt, mailsender, store)
	server := api.NewServer(store, env, taskClient, mailsender)
	// defer config.CloseDB()
	server.Run(":8080")
}

func StartRedisWorker(redisOpts asynq.RedisClientOpt, mailer *utils.MailSender, store *sqlc.SQLStore) {
	log.Info().Msg("Start Task processor")
	client := worker.NewProcessorTaskClient(redisOpts, mailer, store)
	err := client.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start worker")
	}
}
