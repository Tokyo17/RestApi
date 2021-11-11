package config

import (
	// "github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"restApi/model"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dsn := "root:Ilovereza123@tcp(127.0.0.1:3306)/project?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{}, &model.Products{})
}
