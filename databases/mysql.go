package databases

import (
	"fmt"
	"log"

	"github.com/bananafried525/gogin-web/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	var err error
	dns := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig("database.username"),
		config.GetConfig("database.ip"),
		config.GetConfig("database.port"),
		config.GetConfig("database.dbname"),
	)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection Err: %v", err)
	}
	sqlDB, err := Db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
}
