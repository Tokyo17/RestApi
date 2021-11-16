package product

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restApi/config"
	"restApi/model"
	"strconv"
)

// ============================[QUERY]=========================

func GetProductController(c echo.Context) error {
	var products []model.Products
	// config.DB.Find(&products)
	config.DB.Preload("Type").Find(&products)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all products",
		"product": products,
	})
}

// ============================[CREATE]=========================

func CreateProductController(c echo.Context) error {
	product := model.Products{}
	c.Bind(&product)

	config.DB.Create(&product)
	config.DB.Preload("Type").Find(&product)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create product",
		"product": product,
	})
}

// ============================[DELETE]=========================

func DeleteProductController(c echo.Context) error {
	product := model.Products{}

	id, _ := strconv.Atoi(c.Param("id"))
	var z error

	config.DB.Model(&product).Where("ID= ?", id).Find(&product)

	z = c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success Delete all product",
		"products": product,
	})
	config.DB.Where("ID = ?", id).Delete(&product)
	return z
}

// ============================[UPDATE]=========================

func UpdateProductController(c echo.Context) error {
	product := model.Products{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&product)
	config.DB.Model(&product).Where("ID= ?", id).Updates(product)
	config.DB.Model(&product).Where("ID= ?", id).Find(&product)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update product",
		"product": product,
	})
}
