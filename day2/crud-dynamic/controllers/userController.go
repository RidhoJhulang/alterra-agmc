package controllers

import (
	"crud-dynamic/entities"
	"crud-dynamic/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserControllers(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserController(c echo.Context) error {
	//get data create user
	u := new(entities.User)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//change into DTO
	userDTO := entities.ToUserDTO(u)

	//change to models
	userModels := entities.AssembUserDTO(userDTO)

	//save user
	if err := database.SaveUser(userModels); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  userModels,
	})
}

func UpdateUserController(c echo.Context) error {
	//get data create user
	u := new(entities.User)

	//get id
	id := c.Param("id")
	idConvert, _ := strconv.Atoi(id)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//change into DTO
	userDTO := entities.ToUserDTO(u)

	//change to models
	userModels := entities.AssembUserDTO(userDTO)

	//save user
	if err := database.UpdateUsers(userModels, idConvert); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     idConvert,
		"users":  userModels,
	})
}

func GetUserById(c echo.Context) error {

	//get id
	id := c.Param("id")
	idConvert, _ := strconv.Atoi(id)

	//get user
	getUser, err := database.GetUserById(idConvert)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  getUser,
	})
}

func DeleteUserById(c echo.Context) error {

	//get id
	id := c.Param("id")
	idConvert, _ := strconv.Atoi(id)

	//get user
	err := database.DeleteUser(idConvert)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "deleted user successfully",
	})
}
