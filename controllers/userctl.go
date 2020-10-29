package controllers

import (
	"log"

	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"github.com/bananafried525/gogin-web/models/request"
	"github.com/bananafried525/gogin-web/models/response"
	"github.com/bananafried525/gogin-web/services"
	"github.com/bananafried525/gogin-web/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user request.User
	var invalidString string
	var res response.ResultResponse

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		invalidString = "Missing or invalid"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40300"}
		c.JSON(403, res.FmtResponse())
		return
	}

	var newUser gormmodels.User
	newUser.UserName = user.UserName
	newUser.Email = user.Email
	newUser.Password = user.Password
	newUser.RoleID = 2

	if !utils.IsEmail(newUser.Email) {
		invalidString = "Missing or invalid"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40300"}
		c.JSON(403, res.FmtResponse())
		return
	}

	newUser = services.CreateUser(newUser)
	res = response.ResultResponse{ReponseMessage: "", ResponseCode: "20000", ResultData: newUser}
	c.JSON(200, res.FmtResponse())

}

func FindUser(c *gin.Context) {
	var res response.ResultResponse
	var invalidString string

	query := c.Request.URL.Query()
	userId := query.Get("userId")
	newUser := services.GetUser(userId)

	if len(newUser) == 0 {
		invalidString = "Data Notfound"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40401"}
		c.JSON(404, res.FmtResponse())
		return
	}
	res = response.ResultResponse{ReponseMessage: "", ResponseCode: "20000", ResultData: newUser}
	c.JSON(200, res.FmtResponse())
}
