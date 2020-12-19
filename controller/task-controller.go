package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func ListTasks(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
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
