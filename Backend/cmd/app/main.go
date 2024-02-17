package main

import (
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/routes"
	"music-app-backend/internal/app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.InitDB()
	utils.ConfigMail()
	routes.UserRouter(r)
	defer config.CloseDB()
	r.Run(":8080")
}
