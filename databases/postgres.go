package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/bananafried525/gogin-web/config"
	"github.com/bananafried525/gogin-web/databases/gormmodels"
	"github.com/bananafried525/gogin-web/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var host interface{}
var user interface{}
var password interface{}
var dbname interface{}
var port interface{}
var dns string

func ConnectPsqlDb() {
	var err error
	getDbEnvironment()
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	utils.HandleError(err)

	db, err := DB.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		log.Println(err)
	} else {
		log.Println(utils.ConnectPg)
		if !DB.Migrator().HasTable(&gormmodels.User{}) {
			DB.Migrator().CreateTable(&gormmodels.User{})
			log.Println("Created table")

		} else if !DB.Migrator().HasTable(&gormmodels.Role{}) {
			DB.Migrator().CreateTable(&gormmodels.Role{})
			DB.Create(&gormmodels.Role{ID: 1, Role: "ADMIN"})
			DB.Create(&gormmodels.Role{ID: 2, Role: "GUEST"})
		}
	}
}

func getDbEnvironment() {
	if os.Getenv("NODE_ENV") == "docker" {
		host = os.Getenv("pgHost")
		user = os.Getenv("pgUsername")
		password = os.Getenv("pgPassword")
		dbname = os.Getenv("pgDBName")
		port = os.Getenv("pgPost")
	} else {
		host = config.GetConfig("database.psql.ip")
		user = config.GetConfig("database.psql.username")
		password = config.GetConfig("database.psql.password")
		dbname = config.GetConfig("database.psql.dbname")
		port = config.GetConfig("database.psql.port")
	}
	dns = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port,
	)
}
