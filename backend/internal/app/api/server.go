package api

import (
	"music-app-backend/internal/app/helper"
	"music-app-backend/internal/app/utils"
	"music-app-backend/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store       *sqlc.Queries
	router      *gin.Engine
	mailsender  *utils.MailSender
	config      *utils.Config
	token_maker *helper.Token
}

func NewServer(store *sqlc.Queries, config *utils.Config) *Server {

	r := gin.Default()
	server := &Server{
		store:       store,
		config:      config,
		token_maker: helper.NewTokenMaker(config.JwtSecretKey),
	}
	server.mailsender = utils.NewMailSender(server.config)
	server.router = r
	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	s.UserRouter()
}

func (s *Server) Run(address string) {
	s.router.Run(address)
}
