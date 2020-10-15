package controllers

import (
	"log"
	"time"

	"github.com/bananafried525/gogin-web/models/request"
	"github.com/gin-gonic/gin"
)

func GetUserName(c *gin.Context) {
	var user request.User
	if c.ShouldBind(&user) == nil {
		log.Println(&user)
	}
	var newUser map[string]request.User
	newUser = make(map[string]request.User)
	newUser["LUL"] = user
	log.Println(user.IsEmpty())
	c.SecureJSON(200, gin.H{
		"result": newUser,
	})
	return
}

func ValidateUser(c *gin.Context) {
	var user request.User
	if err := c.ShouldBind(&user); err != nil {
		c.SecureJSON(403, gin.H{
			"result": "",
			"msg":    "Missing or invalid",
		})
		return
	} else {
		time.Sleep(time.Second)
		c.SecureJSON(200, gin.H{
			"result": user,
		})
		return
	}
}

func SetUserName(c *gin.Context) {
	var user request.User
	if err := c.ShouldBind(&user); err != nil {
		c.SecureJSON(403, gin.H{
			"result": "",
			"msg":    "Missing or invalid",
		})
		return
	} else {
		user.SetUserName("New User")
		c.SecureJSON(200, gin.H{
			"result": user,
		})
		return
	}
}
