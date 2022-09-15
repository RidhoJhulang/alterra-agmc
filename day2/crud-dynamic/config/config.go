package config

import (
	"crud-dynamic/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var e error

	dsn := "host=127.0.0.1 user=postgres password=asdasdasd dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
