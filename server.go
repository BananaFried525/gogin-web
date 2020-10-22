package main

import (
	"fmt"

	"github.com/bananafried525/gogin-web/config"
	database "github.com/bananafried525/gogin-web/databases"
	"github.com/bananafried525/gogin-web/routes"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	// start web
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	connectDb()
	c := cron.New()
	c.AddFunc("@every 0h15m", logEvery5sec)
	c.Start()

	routes.Routes(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"result": "PAGE_NOT_FOUND", "msg": "Page not found"})
	})
	r.Run(fmt.Sprintf(": %v", config.GetConfig("app.port")))
}

func connectDb() {
	database.ConnectDb()
}

func logEvery5sec() {
	// log.Print("cron")
	// sqlDB, err := database.Db.DB()
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = sqlDB.Ping()
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println("health")
	// }
}
