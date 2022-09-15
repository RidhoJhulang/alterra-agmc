package routes

import (
	"crud-static/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Route Books
	g := e.Group("/books")
	g.GET("", controllers.GetBooksController)
	g.GET("/:id", controllers.GetBookByIdController)
	g.POST("", controllers.CreateBookController)
	g.PUT("/:id", controllers.UpdateBookController)
	g.DELETE("/:id", controllers.DeleteBookControllers)

	return e
}
