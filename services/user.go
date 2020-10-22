package services

import (
	"github.com/bananafried525/gogin-web/databases"
	"github.com/bananafried525/gogin-web/databases/gormmodels"
)

func GetUser() interface{} {
	var users []gormmodels.User
	databases.Db.Find(&users)
	return users
}

func CreateUser(newUser gormmodels.User) gormmodels.User {
	user := newUser
	databases.Db.Create(&user)
	return user
}
