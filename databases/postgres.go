package databases

import (
	"fmt"
	"log"

	"github.com/bananafried525/gogin-web/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var psqlDb *gorm.DB

func ConnectPsqlDb() {
	var err error
	dns := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai ",
		config.GetConfig("database.username"),
		config.GetConfig("database.ip"),
		config.GetConfig("database.port"),
		config.GetConfig("database.dbname"),
	)
	Db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection Err: %v", err)
	}
	sqlDB, err := Db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
}
