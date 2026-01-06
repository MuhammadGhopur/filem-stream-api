package repositories

import (
	"anime-stream-api/config"
	"anime-stream-api/models"
)

// ambil semua film
func GetAllFilms(films *[]models.Film) error {
	return config.DB.Find(films).Error
}

// tambah film (admin)
func CreateFilm(film *models.Film) error {
	return config.DB.Create(film).Error
}
