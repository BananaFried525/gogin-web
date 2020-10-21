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

func CreateUser() interface{} {
	user := gormmodels.User{UserName: "test"}
	databases.Db.Create(&user)
	return user
}
