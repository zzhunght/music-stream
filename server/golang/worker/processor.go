package worker

import (
	"music-app-backend/internal/app/utils"
	"music-app-backend/sqlc"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type ProcessorRedisTasks struct {
	client *asynq.Server
	mailer *utils.MailSender
	store  *sqlc.SQLStore
	rdb    *redis.Client
}

func NewProcessorTaskClient(
	opts asynq.RedisClientOpt,
	mailer *utils.MailSender,
	store *sqlc.SQLStore,
	rdb *redis.Client,
) *ProcessorRedisTasks {
	client := asynq.NewServer(opts, asynq.Config{})
	return &ProcessorRedisTasks{
		client: client,
		mailer: mailer,
		store:  store,
		rdb:    rdb,
	}
}

func (process *ProcessorRedisTasks) Start() error {

	mux := asynq.NewServeMux()

	mux.HandleFunc(TypeEmailDeliveryTask, process.HandleEmailDeliveryTask)
	mux.HandleFunc(TypeEmailDeliveryForgetPasswordRequest, process.HandleEmailDeliveryForgetPasswordRequestTask)
	mux.HandleFunc(TypeEmailDeliveryNewPassword, process.HandleEmailDeliveryNewtPasswordTask)

	return process.client.Start(mux)
}
