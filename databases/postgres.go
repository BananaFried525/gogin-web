package databases

import (
	"fmt"
	"log"

	"github.com/bananafried525/gogin-web/config"
	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var psqlDb *gorm.DB

func ConnectPsqlDb() {
	var err error
	// sslmode=disable TimeZone=Asia/Shanghai
	dns := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=asia/bangkok",
		config.GetConfig("database.psql.username"),
		config.GetConfig("database.psql.password"),
		config.GetConfig("database.psql.dbname"),
		config.GetConfig("database.psql.port"),
	)
	psqlDb, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection Err: %v", err)
	}
	sqlDB, err := psqlDb.DB()

	if !psqlDb.Migrator().HasTable(&gormmodels.User{}) {
		psqlDb.Migrator().CreateTable(&gormmodels.User{})
		log.Println("Created table 'users'")
	} else {
		psqlDb.Migrator().DropTable(&gormmodels.User{})
		psqlDb.Migrator().CreateTable(&gormmodels.User{})
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	defer sqlDB.Close()
	if err := sqlDB.Ping(); err != nil {
		log.Fatalln(err)
		return
	} else {
		log.Println("Hello Postgres")
	}
}
