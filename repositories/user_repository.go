package repositories

import (
	"anime-stream-api/config"
	"anime-stream-api/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).First(user).Error
}

func GetUserByID(user *models.User, id uint) error {
	return config.DB.Where("id = ?", id).First(user).Error
}

func UpdateUser(user models.User) error {
	return config.DB.Save(user).Error
}
