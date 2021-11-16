package routes

import (
	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"restApi/contants"
	"restApi/controller/login"
	"restApi/controller/product"
	"restApi/controller/type"
	"restApi/controller/user"
	"restApi/middlewares"
)

func New() *echo.Echo {

	e := echo.New()

	eNews := e.Group("/users")

	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte(contants.SECRET_JWT),
	}

	eNews.Use(middleware.JWTWithConfig(config))

	eNews.GET("", user.GetUserController)
	e.GET("/users/:id", user.FilterUserController)
	e.GET("/users/deleted", user.FilterDeletedUserController)
	e.POST("/users", user.CreateUserController)
	e.DELETE("/users/:id", user.DeleteUserController)
	e.PUT("/users/:id", user.UpdateUserController)

	e.GET("/products", product.GetProductController)
	e.POST("/products", product.CreateProductController)
	e.DELETE("/products/:id", product.DeleteProductController)
	e.PUT("/products/:id", product.UpdateProductController)

	e.GET("/types", typez.GetTypeController)
	e.POST("/types", typez.CreateTypeController)
	e.DELETE("/types/:id", typez.DeleteTypeController)
	// e.PUT("/products/:id", product.UpdateProductController)

	e.POST("/login", login.LoginController)
	e.Start(":8000")

	return e
}
