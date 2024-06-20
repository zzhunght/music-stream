package api

import (
	"music-app/authentication-services/internal/config"
	"music-app/authentication-services/internal/handler"
	"music-app/authentication-services/internal/helper"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow requests from any origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow the Authorization header
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")

		// Allow GET, POST, OPTIONS methods
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		// If it's an OPTIONS request, we're handling a preflight request.
		// So we don't need to execute the actual handler.
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// Call the next handler
		c.Next()
	}
}

type Server struct {
	router *gin.Engine
	config *config.Config
}

func NewServer() *Server {
	router := gin.Default()

	// Load configuration
	config, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	// Configure CORS
	cor_config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(cor_config))

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
	v1.Use(CORSMiddleware())
	{
		v1.GET("/authentication", handler.Authentication)
	}
}

func (s *Server) Run() error {
	log.Info().Msg("Starting authentication server")
	return s.router.Run(s.config.ServerPort)
}
