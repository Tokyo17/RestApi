package main

import (
	"github.com/labstack/echo/v4"
	"restApi/config"
	"restApi/controller"
	"restApi/model"
)

var Datas []model.User

func main() {
	config.InitDB()
	Datas = []model.User{}

	e := echo.New()

	e.GET("/users", controller.GetUserController)
	e.GET("/users/:id", controller.FilterUserController)
	e.GET("/users/deleted", controller.FilterDeletedUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	e.Start(":8000")
}
