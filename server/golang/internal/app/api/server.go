package api

import (
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/helper"
	"music-app-backend/internal/app/router"
	"music-app-backend/internal/app/utils"
	"music-app-backend/message"
	"music-app-backend/sqlc"
	"music-app-backend/worker"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	store         *sqlc.SQLStore
	router        *gin.Engine
	mailsender    *utils.MailSender
	config        *config.Config
	token_maker   *helper.Token
	task_client   *worker.DeliveryTaskClient
	message_queue *message.RabbitMQProvider
	rdb           *redis.Client
}

func NewServer(
	store *sqlc.SQLStore,
	config *config.Config,
	task_client *worker.DeliveryTaskClient,
	mailsender *utils.MailSender,
	message_queue *message.RabbitMQProvider,
	rdb *redis.Client,
) *Server {
	r := gin.Default()
	cor_config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	r.Use(cors.New(cor_config))
	server := &Server{
		store:         store,
		config:        config,
		token_maker:   helper.NewTokenMaker(config.JwtSecretKey),
		task_client:   task_client,
		message_queue: message_queue,
		rdb:           rdb,
	}
	server.mailsender = mailsender
	server.router = r
	router.SetupRouter(server.router, *server.store, task_client, message_queue, rdb, server.token_maker)
	return server
}

func (s *Server) Run(address string) {
	s.router.Run(address)
}
