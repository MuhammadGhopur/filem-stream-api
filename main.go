package main

import (
	"anime-stream-api/config"
	"anime-stream-api/models"
	"anime-stream-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":3000")
}
