package services

import (
	"anime-stream-api/models"
	"anime-stream-api/repositories"
	"anime-stream-api/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func RegisterUser(name, email, password, role string) error {

	var user models.User

	err := repositories.GetUserByEmail(&user, email)
	if err == nil {
		return errors.New("user already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	newUser := models.User{
		Username: name,
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}

	return repositories.CreateUser(&newUser)
}

func Login(email, password string) (string, error) {
	var user models.User

	fmt.Println("HASH DB     :", user.Password)

	err := repositories.GetUserByEmail(&user, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("email atau password salah")
		}
		return "", err
	}

	err = utils.VerifyPassword(password, user.Password)
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
