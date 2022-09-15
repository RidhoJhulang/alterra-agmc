package routes

import (
	"crud-dynamic/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	u := e.Group("/users")
	u.GET("", controllers.GetUserControllers)
	u.POST("", controllers.CreateUserController)
	u.PUT("/:id", controllers.UpdateUserController)
	u.GET("/:id", controllers.GetUserById)
	u.DELETE("/:id", controllers.DeleteUserById)

	return e
}
