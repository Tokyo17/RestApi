package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restApi/config"
	"restApi/model"
	"strconv"
)

var users []model.User
var user model.User

// ============================[QUERY]=========================

func GetUserController(c echo.Context) error {
	config.DB.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func FilterUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Model(&users).Where("ID= ?", id).Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func FilterDeletedUserController(c echo.Context) error {

	config.DB.Unscoped().Where("Deleted IS NOT NULL").Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get deleted user",
		"users":   users,
	})
}

// ============================[CREATE]=========================

func CreateUserController(c echo.Context) error {
	// name := c.FormValue("name")
	// user.Name = name
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

	id, _ := strconv.Atoi(c.Param("id"))
	var z error

	config.DB.Model(&user).Where("ID= ?", id).Find(&user)

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
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&user)
	config.DB.Model(&user).Where("ID= ?", id).Updates(user)
	config.DB.Model(&user).Where("ID= ?", id).Find(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update users",
		"users":   user,
	})
}
