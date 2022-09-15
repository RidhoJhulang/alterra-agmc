package controllers

import (
	"crud-static/lib/database"
	"crud-static/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func GetBookByIdController(c echo.Context) error {

	id := c.Param("id")

	idConvert, _ := strconv.Atoi(id)

	//get data from books
	books, err := database.GetBookById(idConvert)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func CreateBookController(c echo.Context) error {
	books := models.Books{}
	c.Bind(&books)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func UpdateBookController(c echo.Context) error {

	id := c.Param("id")
	idConvert, _ := strconv.Atoi(id)

	u := new(models.Books)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//change into DTO
	createBook := models.ToBooksDTO(u)

	//from DTO to models
	bookModels := models.AssembBooksDTO(createBook)

	//get data from books
	books, err := database.UpdateBookById(idConvert, bookModels)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "data book has been updated",
		"data":   books,
	})

}

func DeleteBookControllers(c echo.Context) error {

	id := c.Param("id")

	idConvert, _ := strconv.Atoi(id)

	//get data from books
	_, err := database.GetBookById(idConvert)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "book has been deleted",
	})
}
