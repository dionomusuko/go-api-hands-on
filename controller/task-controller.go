package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dionomusuko/go-api-hands-on/database"
)

func ListTasks(c echo.Context) error {
	tasks, err := database.List()
	if err != nil {
		return c.String(http.StatusInternalServerError, "status internal")
	}
	log.Println(tasks)
	return c.JSON(http.StatusOK, tasks)
}

func FindById(c echo.Context) error {
	return nil
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
