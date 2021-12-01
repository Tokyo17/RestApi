package routes

import (
	"final/controllers/admins"
	"final/controllers/products"
	"final/controllers/transactions"
	"final/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController        users.UserController
	TransactionController transactions.TransactionController
	ProductController     products.ProductController
	AdminController       admins.AdminController
	JWTMiddleware         middleware.JWTConfig
	JWTAdmin              middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("users/logins", cl.UserController.Login)
	e.GET("users/details/:id", cl.UserController.Details, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.POST("users/registers", cl.UserController.Register)

	e.GET("payments", cl.TransactionController.GetPM)
	e.POST("payments", cl.TransactionController.AddPM, middleware.JWTWithConfig(cl.JWTAdmin))

	e.POST("shipments", cl.TransactionController.AddShipment, middleware.JWTWithConfig(cl.JWTAdmin))
	e.GET("shipments", cl.TransactionController.GetShipment)

	e.POST("transactions", cl.TransactionController.Add, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("transactions", cl.TransactionController.DetailSC, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.POST("transactions/pay", cl.TransactionController.Pay, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("transactions/details", cl.TransactionController.GetTransDetail, middleware.JWTWithConfig(cl.JWTMiddleware))

	e.POST("products/types", cl.ProductController.UploadType, middleware.JWTWithConfig(cl.JWTAdmin))

	e.GET("products", cl.ProductController.Get)
	e.POST("products", cl.ProductController.UploadProduct, middleware.JWTWithConfig(cl.JWTAdmin))
	e.GET("products/:id", cl.ProductController.FilterByType)
	e.PUT("products", cl.ProductController.UpdateProduct, middleware.JWTWithConfig(cl.JWTAdmin))

	e.POST("admins/registers", cl.AdminController.Register)
	e.POST("admins/logins", cl.AdminController.Login)

}
