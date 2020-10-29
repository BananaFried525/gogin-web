package databases

import "gorm.io/gorm"

var DB *gorm.DB

func ConnectDb() {
	ConnectPsqlDb()
}
