package users_helper

import (
	"github.com/zidanhafiz/go-restapi/models"
)

func GetAllUsers() (*[]models.User, error) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	result := models.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUsernameOrEmail(username string, email string) (*models.User, error) {
	var user models.User
	result := models.DB.Where("username = ? OR email = ?", username, email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUserByID(id string, data models.User) error {
	var user models.User
	result := models.DB.Model(&user).Where("id = ?", id).Updates(data).First(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateUser(username string, email string, password string) (*models.User, error) {
	user := &models.User{Username: username, Email: email, Password: password}
	result := models.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func DeleteUserByID(userId string) error {
	var user models.User
	result := models.DB.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
