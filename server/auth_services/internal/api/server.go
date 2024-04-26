package api

import (
	"music-app/authentication-services/internal/config"
	"music-app/authentication-services/internal/handler"
	"music-app/authentication-services/internal/helper"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func NewServer() *Server {
	router := gin.Default()
	config, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	token_maker := helper.NewTokenMaker(config.JwtSecretKey)
	handler := handler.NewHandler(config, token_maker)
	server := &Server{
		router: router,
		config: config,
	}
	server.SetUpRouter(handler)
	return server
}

func (s *Server) SetUpRouter(handler *handler.Handler) {
	v1 := s.router.Group("/auth")
	{
		v1.GET("/authentication", handler.Authentication)
	}

}

func (s *Server) Run() error {
	log.Info().Msg("Starting authentication server")
	return s.router.Run(s.config.ServerPort)
}
