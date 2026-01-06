package controllers

import (
	"anime-stream-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /films
func GetFilms(c *gin.Context) {
	films, err := services.GetFilms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, films)
}

// POST /films (admin)
func CreateFilm(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		VideoURL    string `json:"video_url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateFilm(req.Title, req.Description, req.VideoURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "film berhasil ditambahkan"})
}
