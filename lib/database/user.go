package database

import (
	"project/config"
	"project/models"
)

func GetAllUsers(interface{} error) {
	var users []models.Users
	
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}


func CreateUser(user models.User) error {
	err := config.DB.Save(&user).Error
	if err != nil {
		return nill, err
	}
	return err
}
