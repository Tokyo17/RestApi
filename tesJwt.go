// package main

// import (
// 	"fmt"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	// "github.com/jinzhu/gorm"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"net/http"
// 	"restApi/contants"
// 	"restApi/middlewares"
// 	// "strconv"
// 	"restApi/controller/login"
// 	"time"
// )

// var pln = fmt.Println
// var DB *gorm.DB

// type User struct {
// 	ID        int
// 	Name      string
// 	Password  string
// 	Token     string
// 	CreatedAt time.Time
// 	DeletedAt gorm.DeletedAt
// }

// func InitDB() {
// 	var err error

// 	dsn := "root:Ilovereza123@tcp(127.0.0.1:3306)/coba?charset=utf8mb4&parseTime=True&loc=Local"
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}
// 	DB.AutoMigrate(&User{})

// }
// func main() {
// 	InitDB()
// 	e := echo.New()

// 	// e.GET("/users", GetNewsController)

// 	eNews := e.Group("/users")
// 	config := middleware.JWTConfig{
// 		Claims:     &middlewares.JwtCustomClaims{},
// 		SigningKey: []byte(contants.SECRET_JWT),
// 	}

// 	eNews.Use(middleware.JWTWithConfig(config))
// 	eNews.GET("", GetUserDetailControllers)

// 	e.POST("/login", login.LoginController)
// 	e.Start(":8000")

// }

// func GetNewsController(c echo.Context) error {
// 	var user []User

// 	DB.Model(&User{}).Preload("Payments").Find(&user)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":    http.StatusOK,
// 		"message": "success get all news",
// 		"users":   user,
// 	})
// }

// // func LoginController(c echo.Context) error {
// // 	user := User{}
// // 	c.Bind(&user)

// // 	DB.Where("email = ? AND password = ?", user.Name, user.Password).First(&user)
// // 	token, _ := middlewares.CreateToken(int(user.ID))
// // 	return c.JSON(http.StatusOK, map[string]interface{}{
// // 		"code":    http.StatusOK,
// // 		"message": "success get all news",
// // 		"TOKEN":   token,
// // 	})
// // }

// func GetUserDetailControllers(c echo.Context) error {
// 	var users User
// 	DB.Find(&users)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":    http.StatusOK,
// 		"message": "success get all news",
// 		"users":   users,
// 	})
// }
