package typez

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restApi/config"
	"restApi/model"
	"strconv"
)

// ============================[QUERY]=========================

func GetTypeController(c echo.Context) error {
	var typez []model.Type
	config.DB.Find(&typez)
	// config.DB.Preload("Type").Find(&types)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all type",
		"product": typez,
	})
}

// ============================[CREATE]=========================

func CreateTypeController(c echo.Context) error {
	typez := model.Type{}
	c.Bind(&typez)

	config.DB.Create(&typez)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create type",
		"product": typez,
	})
}

// ============================[DELETE]=========================

func DeleteTypeController(c echo.Context) error {
	typez := model.Type{}

	id, _ := strconv.Atoi(c.Param("id"))
	var z error

	config.DB.Model(&typez).Where("ID= ?", id).Find(&typez)

	z = c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success Delete  type",
		"products": typez,
	})
	config.DB.Where("ID = ?", id).Delete(&typez)
	return z
}

// // ============================[UPDATE]=========================

func UpdateProductController(c echo.Context) error {
	typez := model.Type{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&typez)
	config.DB.Model(&typez).Where("ID= ?", id).Updates(typez)
	config.DB.Model(&typez).Where("ID= ?", id).Find(&typez)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update product",
		"product": typez,
	})
}
