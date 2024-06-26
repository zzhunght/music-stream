package api

import (
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/helper"
	"music-app-backend/internal/app/utils"
	"music-app-backend/message"
	"music-app-backend/pkg/middleware"
	"music-app-backend/sqlc"
	"music-app-backend/worker"
	"net/http"

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
	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {

	v1 := s.router.Group("/api/v1")
	{
		v1.GET("health-check", healthCheck)
	}
	v1.Use(middleware.CORSMiddleware())
	s.UserRouter(v1)
	s.AdminRouter(v1)
	s.PublicRouter(v1)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (s *Server) Run(address string) {
	s.router.Run(address)
}

func SuccessResponse(data interface{}, message string) gin.H {
	return gin.H{
		"data":    data,
		"message": message,
	}
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
