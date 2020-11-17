package services

import (
	"log"

	"github.com/bananafried525/gogin-web/databases"
	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"github.com/bananafried525/gogin-web/utils"
	"github.com/dgrijalva/jwt-go"
)

func Login(user *gormmodels.User) bool {
	password := user.Password
	databases.DB.Preload("Role").Where("user_name = ?", user.UserName).Find(&user)
	if utils.CheckPasswordHash(password, user.Password) {
		return true
	}
	return false
}

func EncodeJwt(auth gormmodels.User) string {
	mySecret := []byte("ohmmho")
	tokenE := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": auth.UserName,
		"role":     auth.Role.Role,
	})
	tokenString, err := tokenE.SignedString(mySecret)
	if err != nil {
		log.Println(err)
		return "error"
	}
	return tokenString
}
