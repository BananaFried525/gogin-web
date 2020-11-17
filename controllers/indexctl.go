package controllers

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"github.com/bananafried525/gogin-web/models/request"
	"github.com/bananafried525/gogin-web/models/response"
	"github.com/bananafried525/gogin-web/services"
	"github.com/bananafried525/gogin-web/utils"
)

func RequestJsonHolder(c *gin.Context) {
	c.JSON(200, "Hello Go")
	return
}

func Test(c *gin.Context) {
	userName, _ := c.GetQuery("userName")
	userEmail, _ := c.GetQuery("userEmail")
	user := gormmodels.NewUser()
	user.UserName = userName
	user.Email = userEmail
	services.FindUser(user)
	if user.ID == 0 {
		c.JSON(403, nil)
		return
	}
	if utils.CheckPasswordHash("asdasdasd", user.Password) {
		log.Println("pass")
	}
	resUser := response.UserResponse{}
	resUser.New(*user)
	c.JSON(200, resUser)
	return
}

func Login(c *gin.Context) {
	log.Println(c.Get("G"))
	var user request.User
	var invalidString string
	// var err error
	var res response.ResultResponse

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		invalidString = "Missing or invalid"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40300"}
		c.JSON(403, res.FmtResponse())
		return
	}
	auth := gormmodels.User{}
	auth.UserName = user.UserName
	auth.Password = user.Password
	if services.Login(&auth) {
		res = response.ResultResponse{ReponseMessage: "login success", ResponseCode: "20000"}
		authJWT := services.EncodeJwt(auth)
		log.Println(authJWT)
		c.Header("auth", authJWT)
		c.JSON(201, res.FmtResponse())
		return
	} else {
		res = response.ResultResponse{ReponseMessage: "", ResponseCode: "40100"}
		c.JSON(401, res.FmtResponse())
		return
	}

}
