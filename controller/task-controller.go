package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/dionomusuko/go-api-hands-on/database"
)

func ListTasks(c echo.Context) error {
	tasks, err := database.List()
	if err != nil {
		return c.String(http.StatusInternalServerError, "status internal")
	}
	return c.JSON(http.StatusOK, tasks)
}

func FindById(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	log.Println(id)
	task, err := database.FindById(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, task)
}

func Create(c echo.Context) error {
	return nil
}

func Delete(c echo.Context) error {
	return nil
}

func Update(c echo.Context) error {
	return nil
}
