package databases

import (
	"fmt"
	"log"

	"github.com/bananafried525/gogin-web/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var Db *gorm.DB

func ConnectMySqlDb() {
	var err error
	dns := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig("database.username"),
		config.GetConfig("database.ip"),
		config.GetConfig("database.port"),
		config.GetConfig("database.dbname"),
	)
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection Err: %v", err)
	}
	db, err := DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
		return
	} else {
		log.Println("Hello Postgres")
	}
}
