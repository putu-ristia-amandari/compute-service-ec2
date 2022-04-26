package routes

import (
	"project/controllers"

	"github.com/labstack/echo"
)

func New() e *echo.Echo {
	e := echo.New

	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUserController)
	

	return e
}