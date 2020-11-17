package routes

import (
	"log"

	"github.com/bananafried525/gogin-web/controllers"
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
		user.POST("/createuser", controllers.CreateUser)
		user.GET("/getusers", controllers.FindUser)
		user.DELETE("/deleteusser", controllers.DeleteUser)
	}
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("")
	}

}
