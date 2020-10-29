package routes

import (
	"github.com/bananafried525/gogin-web/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	index := route.Group("/")
	{
		index.GET("/", controllers.RequestJsonHolder)
	}
	user := route.Group("/user")
	{
		user.POST("/createuser", controllers.CreateUser)
		user.GET("/getusers", controllers.FindUser)
		// user.GET("/getuser")
	}
}
