package services

import (
	"anime-stream-api/models"
	"anime-stream-api/repositories"
)

// user & admin
func GetFilms() ([]models.Film, error) {
	var films []models.Film
	err := repositories.GetAllFilms(&films)
	return films, err
}

// admin only
func CreateFilm(title, desc, videoURL string) error {
	film := models.Film{
		Title:       title,
		Description: desc,
		VideoURL:    videoURL,
	}
	return repositories.CreateFilm(&film)
}
