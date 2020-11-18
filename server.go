package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bananafried525/gogin-web/config"
	"github.com/bananafried525/gogin-web/databases"
	"github.com/bananafried525/gogin-web/routes"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	env := os.Getenv("NODE_ENV")
	var port interface{}
	if env == "docker" {
		port = os.Getenv("PORT")
	} else {
		port = config.GetPort()
	}
	log.Println(os.Getenv("test"))
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
	r.Run(fmt.Sprintf(":%v", port))
}

func connectDb() {
	databases.ConnectDb()
}

func logEvery5sec() {
	sqlDB, err := databases.DB.DB()
	if err != nil {
		log.Println(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Println(err)
		sqlDB.Close()
		databases.ConnectDb()
	} else {
		log.Println("db connected")
	}
}

func heyJwt() {
	mySecret := []byte("ohmmho")
	tokenE := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"dd":   "qwe",
		"dasd": "123123",
	})
	tokenString, err := tokenE.SignedString(mySecret)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(tokenString)

	tokenD, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySecret, nil
	})
	claims, ok := tokenD.Claims.(jwt.MapClaims)
	if ok && tokenD.Valid {
		fmt.Println(claims["dd"], claims["dasd"])
	} else {
		fmt.Println(err)
	}
}
