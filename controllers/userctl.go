package controllers

import (
	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"github.com/bananafried525/gogin-web/models/request"
	"github.com/bananafried525/gogin-web/models/response"
	"github.com/bananafried525/gogin-web/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user request.User
	var invalidString string
	var res response.ResultResponse

	if err := c.ShouldBindJSON(&user); err != nil {
		invalidString = "Missing or invalid"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40300"}
		c.JSON(403, res.FmtResponse())
		return
	}
	var newUser gormmodels.User
	newUser.UserName = user.UserName
	newUser = services.CreateUser(newUser)
	res = response.ResultResponse{ReponseMessage: "", ResponseCode: "20000", ResultData: newUser}
	c.JSON(200, res.FmtResponse())

}
