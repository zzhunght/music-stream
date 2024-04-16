package main

import (
	"fmt"
	"music-app-backend/gapi"
	"music-app-backend/internal/app/api"
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/utils"
	"music-app-backend/message"
	"music-app-backend/pb"
	"music-app-backend/sqlc"
	"music-app-backend/worker"
	"net"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	go StartGRPCServer(store, env, taskClient, mailsender, mq, rdb)
	StartHttpServer(store, env, taskClient, mailsender, mq, rdb)
}

func StartHttpServer(
	store *sqlc.SQLStore,
	config *utils.Config,
	task_client *worker.DeliveryTaskClient,
	mailsender *utils.MailSender,
	message_queue *message.RabbitMQProvider,
	rdb *redis.Client,
) {
	server := api.NewServer(store, config, task_client, mailsender, message_queue, rdb)
	// defer config.CloseDB()
	server.Run(":8080")
}

func StartGRPCServer(
	store *sqlc.SQLStore,
	config *utils.Config,
	task_client *worker.DeliveryTaskClient,
	mailsender *utils.MailSender,
	message_queue *message.RabbitMQProvider,
	rdb *redis.Client,
) {
	fmt.Println("Started grpc server ......")
	server := gapi.NewServer(store, config, task_client, mailsender, message_queue, rdb)
	grpcServer := grpc.NewServer()
	pb.RegisterMusicAppServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}
	fmt.Println("Started grpc server at port 9090")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start grpc")
	}
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
