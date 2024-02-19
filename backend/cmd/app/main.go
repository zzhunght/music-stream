package main

import (
	"music-app-backend/internal/app/api"
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/utils"
)

func main() {
	store := config.InitDB()
	utils.ConfigMail()
	server := api.NewServer(store)

	defer config.CloseDB()
	server.Run(":8080")
}
