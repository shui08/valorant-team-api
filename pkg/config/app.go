package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var placeholder string

func Connect() {
	data, err := gorm.Open(mysql.Open(placeholder), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = data
}

func GetDB() *gorm.DB {
	return db
}
