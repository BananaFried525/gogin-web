package controllers

import (
	"log"
	"strings"

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
	var err error
	var res response.ResultResponse

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		invalidString = "Missing or invalid"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40300"}
		c.JSON(403, res.FmtResponse())
		return
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		invalidString = "Missing or invalid"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40300"}
		c.JSON(403, res.FmtResponse())
		return
	}

	var newUser gormmodels.User
	newUser.NewGuest(user, 0)

	if !utils.IsEmail(newUser.Email) {
		invalidString = "Missing or invalid"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40300"}
		c.JSON(403, res.FmtResponse())
		return
	}

	newUser, err = services.CreateUser(newUser)
	if err != nil {
		invalidString = "Error Database"
		if strings.Contains(err.Error(), "23505") {
			res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "50000", ResultData: "duplicate key"}
			c.JSON(500, res.FmtResponse())
			return
		} else {
			res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "50000", ResultData: err.Error()}
			c.JSON(500, res.FmtResponse())
			return
		}

	}

	res = response.ResultResponse{ReponseMessage: "", ResponseCode: "20100", ResultData: "Create Success"}
	c.JSON(201, res.FmtResponse())
	return
}

func FindUser(c *gin.Context) {
	var res response.ResultResponse
	var invalidString string

	query := c.Request.URL.Query()
	userID := query.Get("userId")
	newUser := services.GetUser(userID)

	if len(newUser) == 0 {
		invalidString = "Data Notfound"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40401"}
		c.JSON(404, res.FmtResponse())
		return
	}

	var userRes []response.UserResponse
	for _, user := range newUser {
		u := response.UserResponse{}
		u.New(user)
		userRes = append(userRes, u)
	}
	res = response.ResultResponse{ReponseMessage: "", ResponseCode: "20000", ResultData: userRes}
	c.JSON(200, res.FmtResponse())
}

func DeleteUser(c *gin.Context) {
	var res response.ResultResponse
	var invalidString string

	query := c.Request.URL.Query()
	userID := query.Get("userId")

	if userID == "" {
		invalidString = "Missing or invalid"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40300"}
		c.JSON(403, res.FmtResponse())
		return
	}

	findUser := services.GetUser(userID)
	if len(findUser) == 0 {
		invalidString = "Data Notfound"
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "40401"}
		c.JSON(404, res.FmtResponse())
		return
	}

	log.Println(&findUser)
	err := services.DeleteUser(&findUser[0])
	if err != nil {
		res = response.ResultResponse{ReponseMessage: invalidString, ResponseCode: "50000", ResultData: err.Error()}
		c.JSON(500, res.FmtResponse())
		return
	}
	res = response.ResultResponse{ReponseMessage: "", ResponseCode: "20000"}
	c.JSON(200, res.FmtResponse())

}
