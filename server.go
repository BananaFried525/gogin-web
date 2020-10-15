package main

import (
	"fmt"

	"github.com/bananafried525/gogin-web/config"
	"github.com/bananafried525/gogin-web/database"
	"github.com/bananafried525/gogin-web/routes"
	"github.com/codeskyblue/go-sh"
	"github.com/gin-gonic/gin"
)

func main() {
	// start web
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	connectDb()
	routes.Routes(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"result": "PAGE_NOT_FOUND", "msg": "Page not found"})
	})
	r.Run(fmt.Sprintf(": %v", config.GetConfig("app.port"))) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func runCmd() {
	// shell run .exe file
	sh.Command("./cmd/cmd.exe").Run()
}

func connectDb() {
	database.Db()
}