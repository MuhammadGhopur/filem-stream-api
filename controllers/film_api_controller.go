package controllers

import (
	"anime-stream-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ===============================
// GET /films/search?q=batman
// ===============================
func SearchFilms(c *gin.Context) {

	keyword := c.Query("q")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query q wajib diisi",
		})
		return
	}

	result, err := services.SearchFilms(keyword)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ===============================
// GET /films/:imdb_id
// ===============================
func GetFilmDetail(c *gin.Context) {

	imdbID := c.Param("imdb_id")

	result, err := services.GetFilmDetail(imdbID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
