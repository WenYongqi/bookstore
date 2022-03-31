package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"gorm.io/gorm"
	"log"
)

var db = utils.DB

func CheckUsernameAndPassword(username string, password string) (*model.User, error) {
	var user model.User
	result := db.Where("username = ? and password = ?", username, password).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Println(result.Error)
		return &user, result.Error
	}
	return &user, nil
}

func CheckUsername(username string) (*model.User, error) {
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Println(result.Error)
		return &user, result.Error
	}
	return &user, nil
}

func SaveUser(username string, password string, email string) error {
	user := model.User{
		Username: username,
		Password: password,
		Email: email,
	}
	result := db.Create(&user)
	log.Println("Rows affected:", result.RowsAffected)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return result.Error
}