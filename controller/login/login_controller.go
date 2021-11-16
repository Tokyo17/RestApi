package login

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"restApi/config"
	"restApi/middlewares"
	"restApi/model"
)

func LoginController(c echo.Context) error {
	// user := model.User{}
	var users model.User
	c.Bind(&users)

	if err := config.DB.Where("email = ? AND password = ?", users.Email, users.Password).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Email atau password salah",
		})
	}

	fmt.Println(users.Address, users.ID)

	token, _ := middlewares.CreateToken(int(users.ID))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get token",
		"TOKEN":   token,
	})
}
