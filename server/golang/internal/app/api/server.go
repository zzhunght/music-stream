package api

import (
	"music-app-backend/internal/app/helper"
	"music-app-backend/internal/app/utils"
	"music-app-backend/message"
	"music-app-backend/sqlc"
	"music-app-backend/worker"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	store         *sqlc.SQLStore
	router        *gin.Engine
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
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	v1 := s.router.Group("/api/v1")
	s.UserRouter(v1)
	s.AdminRouter(v1)
	s.PublicRouter(v1)
}

func (s *Server) Run(address string) {
	s.router.Run(address)
}

func SuccessResponse(data any, messgae string) gin.H {
	return gin.H{
		"data":    data,
		"messgae": messgae,
	}
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
