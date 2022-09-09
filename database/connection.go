package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:verysecret@tcp(localhost:33061)/yt_go_auth?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
}
