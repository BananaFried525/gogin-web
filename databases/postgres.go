package databases

import (
	"fmt"
	"log"

	"github.com/bananafried525/gogin-web/config"
	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPsqlDb() {
	var err error
	// sslmode=disable TimeZone=Asia/Shanghai
	dns := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=asia/bangkok",
		config.GetConfig("database.psql.username"),
		config.GetConfig("database.psql.password"),
		config.GetConfig("database.psql.dbname"),
		config.GetConfig("database.psql.port"),
	)
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection Err: %v", err)
	}

	if !DB.Migrator().HasTable(&gormmodels.User{}) || !DB.Migrator().HasTable(&gormmodels.Role{}) {
		DB.Migrator().CreateTable(&gormmodels.Role{})
		DB.Migrator().CreateTable(&gormmodels.User{})
		log.Println("Created table")
	} else {
		DB.Migrator().DropTable(&gormmodels.Role{})
		DB.Migrator().DropTable(&gormmodels.User{})
		log.Println("Deleted")
		DB.Migrator().CreateTable(&gormmodels.Role{})
		DB.Migrator().CreateTable(&gormmodels.User{})
		log.Println("Created table")
		DB.Create(&gormmodels.Role{ID: 1, Role: "ADMIN"})
		DB.Create(&gormmodels.Role{ID: 2, Role: "GUEST"})
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