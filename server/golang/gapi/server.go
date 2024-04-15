package gapi

import (
	"music-app-backend/internal/app/helper"
	"music-app-backend/internal/app/utils"
	"music-app-backend/message"
	"music-app-backend/pb"
	"music-app-backend/sqlc"
	"music-app-backend/worker"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	pb.UnimplementedMusicAppServer
	store         *sqlc.SQLStore
	mailsender    *utils.MailSender
	config        *utils.Config
	token_maker   *helper.Token
	task_client   *worker.DeliveryTaskClient
	message_queue *message.RabbitMQProvider
	rdb           *redis.Client
}

func NewServer(
	store *sqlc.SQLStore,
	config *utils.Config,
	task_client *worker.DeliveryTaskClient,
	mailsender *utils.MailSender,
	message_queue *message.RabbitMQProvider,
	rdb *redis.Client,
) *Server {
	server := &Server{
		store:         store,
		config:        config,
		token_maker:   helper.NewTokenMaker(config.JwtSecretKey),
		task_client:   task_client,
		message_queue: message_queue,
		rdb:           rdb,
	}
	server.mailsender = mailsender
	return server
}

func (s *Server) Run(address string) {

}
