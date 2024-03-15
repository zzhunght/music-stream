package api

import (
	"music-app-backend/internal/app/helper"
	"music-app-backend/internal/app/utils"
	"music-app-backend/sqlc"
	"music-app-backend/worker"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store       *sqlc.SQLStore
	router      *gin.Engine
	mailsender  *utils.MailSender
	config      *utils.Config
	token_maker *helper.Token
	task_client *worker.DeliveryTaskClient
}

func NewServer(
	store *sqlc.SQLStore,
	config *utils.Config,
	task_client *worker.DeliveryTaskClient,
	mailsender *utils.MailSender,
) *Server {
	r := gin.Default()
	server := &Server{
		store:       store,
		config:      config,
		token_maker: helper.NewTokenMaker(config.JwtSecretKey),
		task_client: task_client,
	}
	server.mailsender = mailsender
	server.router = r
	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	s.UserRouter()
	s.AdminRouter()
	s.PublicRouter()
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
