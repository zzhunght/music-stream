package main

import (
	"music-app-backend/internal/app/api"
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/utils"
	"music-app-backend/message"
	"music-app-backend/sqlc"
	"music-app-backend/worker"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
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

	mq := StartRabbitMQ(env)
	defer mq.CloseRabbitMQ()
	taskClient := worker.NewDeliveryTaskClient(redisOpt)
	mailsender := utils.NewMailSender(env)
	rdb := utils.NewRedisClient(env.RedisUrl)
	defer rdb.Close()

	store := config.InitDB(env.DatabaseDestination)

	go StartRedisWorker(redisOpt, mailsender, store, rdb)
	server := api.NewServer(store, env, taskClient, mailsender, mq, rdb)
	// defer config.CloseDB()
	server.Run(":8080")
}

func StartRedisWorker(
	redisOpts asynq.RedisClientOpt,
	mailer *utils.MailSender,
	store *sqlc.SQLStore,
	rdb *redis.Client,
) {
	log.Info().Msg("Start Task processor")
	client := worker.NewProcessorTaskClient(redisOpts, mailer, store, rdb)
	err := client.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start worker")
	}
}

func StartRabbitMQ(config *utils.Config) *message.RabbitMQProvider {
	messageQueue, err := message.NewRabbitMQ(config)
	if err != nil {
		log.Fatal().Err(err).Msg(("can connect to message queue"))
	}
	err = messageQueue.DeclareExchange()
	if err != nil {
		log.Fatal().Err(err).Msg(("can declare exchange"))
	}
	log.Info().Msg("Start RabbitMQ Successfully!")
	return messageQueue
}
