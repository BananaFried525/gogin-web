package routes

import (
	"log"
	"net/http"

	"github.com/bananafried525/gogin-web/controllers"
	"github.com/bananafried525/gogin-web/services"
	"github.com/bananafried525/gogin-web/utils"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	index := route.Group("/")
	{
		index.GET("/", controllers.RequestJsonHolder)
		index.GET("/test", controllers.Test)
		index.POST("/login", middleware(), controllers.Login)
	}
	user := route.Group("/user")
	{
		user.POST("/createuser", getJwt(), controllers.CreateUser)
		user.GET("/getusers", controllers.FindUser)
		user.DELETE("/deleteusser", controllers.DeleteUser)
	}
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("")
		c.Next()
	}

}

func getJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("auth")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}
		claims, ok := services.DecodeJwt(token)
		c.Set("data", claims)
		if (!ok && claims["role"] != "ADMIN") || !utils.CheckJwtExpire(claims["exp"]) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}
		c.Next()
	}
}
