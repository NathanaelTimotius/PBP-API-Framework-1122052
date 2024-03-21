package main

import (
	"latihan_echo/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/users", controllers.GetAllUsersEcho)
	e.POST("/users", controllers.InsertUserEcho)
	e.PUT("/users", controllers.UpdateUserEcho)
	e.DELETE("/users", controllers.DeleteUserEcho)

	e.Logger.Fatal(e.Start(":8888"))
}
