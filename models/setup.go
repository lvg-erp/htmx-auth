package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(os.Getenv("DB_CONN_STR"), &gorm.Config{}))

	if err != nil {
		panic("failed to connect to database")
	}

	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Blog{}, &User)
}
