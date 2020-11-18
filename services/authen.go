package services

import (
	"fmt"
	"time"

	"github.com/bananafried525/gogin-web/databases"
	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"github.com/bananafried525/gogin-web/utils"
	"github.com/dgrijalva/jwt-go"
)

func Login(user *gormmodels.User) bool {
	password := user.Password
	err := databases.DB.Preload("Role").Where("user_name = ?", user.UserName).Find(&user).Error
	utils.HandleError(err)
	if utils.CheckPasswordHash(password, user.Password) {
		return true
	}
	return false
}

func EncodeJwt(auth gormmodels.User) string {
	mySecret := []byte(utils.JWTKEY)
	exp := time.Now().Local().Add(time.Hour).Format(time.RFC3339)
	tokenE := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": auth.UserName,
		"role":     auth.Role.Role,
		"exp":      exp,
	})
	tokenString, err := tokenE.SignedString(mySecret)
	utils.HandleError(err)
	return tokenString
}

func DecodeJwt(tokenString string) (map[string]interface{}, bool) {
	tokenD, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.JWTKEY), nil
	})
	utils.HandleError(err)
	claims, ok := tokenD.Claims.(jwt.MapClaims)
	return claims, ok
}
