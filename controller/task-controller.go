package controller

import (
	"log"

	"github.com/labstack/echo/v4"
)

func ListTasks(c echo.Context) error {
	log.Println("test")
	return c.JSON(200, "ok")
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
