package api

import (
	"music-app-backend/internal/app/utils"
	"music-app-backend/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store      *sqlc.Queries
	router     *gin.Engine
	mailsender *utils.MailSender
}

func NewServer(store *sqlc.Queries) *Server {

	r := gin.Default()
	server := &Server{
		store: store,
	}
	server.mailsender = utils.NewMailSender()
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
