package database

import (
	. "github.com/Samito19/msu-mp-auth-ms/errorHandlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	dsn := "root:toor@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	CheckError(err)
	return db
}
