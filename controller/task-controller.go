package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/dionomusuko/go-api-hands-on/database"
	"github.com/dionomusuko/go-api-hands-on/model"
)

func ListTasks(c echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := database.List(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "status internal")
	}
	return c.JSON(http.StatusOK, tasks)
}

func FindById(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	task, err := database.FindById(ctx, id)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, task)
}

func Create(c echo.Context) error {
	ctx := c.Request().Context()

	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		c.Logger().Error(err)
	}
	if err := database.Store(ctx, task); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusCreated, task)
}

func Delete(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := database.Delete(ctx, id); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.String(http.StatusOK, "task deleted")
}

func Update(c echo.Context) error {
	ctx := c.Request().Context()

	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		c.Logger().Error(err)
	}

	if err := database.Update(ctx, task); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, task)
}
