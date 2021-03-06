package services

import (
	"log"

	"github.com/bananafried525/gogin-web/databases"
	"github.com/bananafried525/gogin-web/databases/gormmodels"
)

func GetUser(id string) []gormmodels.User {
	var users []gormmodels.User
	err := databases.DB.Preload("Role").Find(&users, id).Error
	if err != nil {
		log.Println(err)
	}

	return users
}

func CreateUser(newUser gormmodels.User) (gormmodels.User, error) {
	user := newUser
	err := databases.DB.Create(&user).Error
	return user, err
}

func DeleteUser(user *gormmodels.User) error {
	err := databases.DB.Delete(&user).Error
	return err
}
