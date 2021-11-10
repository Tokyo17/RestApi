package controller

import (
	"github.com/labstack/echo/v4"

	"net/http"
	"restApi/config"
	"restApi/model"
)

// ============================[QUERY]=========================

func GetUserController(c echo.Context) error {
	var users []model.User
	config.DB.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// ============================[CREATE]=========================

func CreateUserController(c echo.Context) error {
	var user = model.User{}

	name := c.FormValue("name")
	user.Name = name

	c.Bind(&user)

	config.DB.Save(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create all users",
		"users":   user,
	})
}

// ============================[DELETE]=========================

func DeleteUserController(c echo.Context) error {
	var user model.User

	c.Bind(&user)

	var z error

	config.DB.Model(&user).Where("ID= ?", user.ID).Find(&user)

	z = c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Delete all users",
		"users":   user,
	})
	config.DB.Where("ID = ?", user.ID).Delete(&user)
	return z

}

// ============================[UPDATE]=========================

func UpdateUserController(c echo.Context) error {
	var user model.User

	c.Bind(&user)
	config.DB.Model(&user).Where("ID= ?", user.ID).Updates(user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create all users",
		"users":   user,
	})
}
