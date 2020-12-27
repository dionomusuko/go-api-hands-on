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

	title := c.QueryParam("title")
	description := c.QueryParam("description")
	t := model.Task{
		Title:       title,
		Description: description,
	}

	if err := database.Store(ctx, t); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.String(http.StatusCreated, "task created")
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

	paramID := c.Param("id")
	title := c.QueryParam("title")
	description := c.QueryParam("description")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	t := model.Task{
		Id:          id,
		Title:       title,
		Description: description,
	}

	if err := database.Update(ctx, t); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.String(http.StatusOK, "task updated")
}
