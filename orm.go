package main

import (
	// "fmt"
	"github.com/labstack/echo/v4"
	// "net/http"
	"restApi/config"
	"restApi/controller"
	"restApi/model"
	// "time"
)

var Datas []model.User

func main() {
	config.InitDB()
	Datas = []model.User{}

	e := echo.New()

	e.GET("/getUser", controller.GetUserController)
	e.POST("/crtUser", controller.CreateUserController)
	e.DELETE("/delUser", controller.DeleteUserController)
	e.PUT("/updUser", controller.UpdateUserController)

	e.Start(":8000")

}
