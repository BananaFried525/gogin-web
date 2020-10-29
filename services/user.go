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

func CreateUser(newUser gormmodels.User) gormmodels.User {
	user := newUser
	databases.DB.Create(&user)
	return user
}
