package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/bananafried525/gogin-web/config"
	_ "github.com/go-sql-driver/mysql"
)

func Db() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s:%s)/%s",
		config.GetConfig("database.username"),
		config.GetConfig("database.ip"),
		config.GetConfig("database.port"),
		config.GetConfig("database.dbname"),
	))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	result, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var id int
		var username string
		result.Scan(&id, &username)
		log.Printf("ID=%d, Title=%s\n", id, username)
	}
}
