package routes

import (
	"anime-stream-api/controllers"
	"anime-stream-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// PUBLIC
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	// JWT REQUIRED
	auth := router.Group("/")
	auth.Use(middlewares.JWTMiddleware())
	{
		auth.GET("/films/search", controllers.SearchFilms)
		auth.GET("/films/:imdb_id", controllers.GetFilmDetail)
	}
}
