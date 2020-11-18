package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"github.com/bananafried525/gogin-web/models/request"
	"github.com/bananafried525/gogin-web/models/response"
	"github.com/bananafried525/gogin-web/services"
	"github.com/bananafried525/gogin-web/utils"
)

func RequestJsonHolder(c *gin.Context) {
	c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
	return
}

func Test(c *gin.Context) {
	add := time.Now().Local().Add(time.Hour)
	c.JSON(200, add.String())
	return
}

func Login(c *gin.Context) {
	var user request.User
	var err error
	var res response.ResultResponse

	if err = c.ShouldBindJSON(&user); err != nil {
		utils.HandleError(err)
		res = response.ResultResponse{ReponseMessage: utils.LoginFail, ResponseCode: "40100"}
		c.JSON(http.StatusUnauthorized, res.FmtResponse())
		return
	}
	auth := gormmodels.User{}
	auth.UserName = user.UserName
	auth.Password = user.Password
	if services.Login(&auth) {
		res = response.ResultResponse{ReponseMessage: utils.LoginSuccess, ResponseCode: "20000"}
		authJWT := services.EncodeJwt(auth)
		c.Header("auth", authJWT)
		c.JSON(http.StatusOK, res.FmtResponse())
	} else {
		res = response.ResultResponse{ReponseMessage: utils.LoginFail, ResponseCode: "40100"}
		c.JSON(http.StatusUnauthorized, res.FmtResponse())
	}

}
