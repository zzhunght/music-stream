package main

import (
	"music-app-backend/internal/app/api"
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/utils"

	"github.com/rs/zerolog/log"
)

func main() {
	env, err := utils.LoadEnv(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	store := config.InitDB(env.DatabaseDestination)
	server := api.NewServer(store, env)
	defer config.CloseDB()
	server.Run(":8080")
}
