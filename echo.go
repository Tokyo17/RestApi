package main

import (
	// "github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type coba struct {
	Nama string
	Age  int
}

var user []coba

func main() {
	user = []coba{{Nama: "Reza", Age: 20}, {Nama: "Ummul", Age: 21}}
	e := echo.New()
	e.GET("/data", allData)
	e.GET("/data/:id", oneData)
	e.POST("/add", add)
	e.Start(":8000")
}
func allData(c echo.Context) error {

	return c.JSON(http.StatusOK, user)
}

func oneData(c echo.Context) error {
	nama, _ := strconv.Atoi(c.Param("id"))
	var a error
	for _, v := range user {
		if nama == v.Age {
			a = c.JSON(http.StatusOK, v)
		}
	}
	return a
}

func add(c echo.Context) error {
	nama := c.FormValue("Nama")
	age := c.FormValue("Age")
	var addUser coba
	addUser.Nama = nama
	addUser.Age, _ = strconv.Atoi(age)
	user = append(user, addUser)

	return c.JSON(http.StatusOK, addUser)
}
