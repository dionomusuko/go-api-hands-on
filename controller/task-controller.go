package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/dionomusuko/go-api-hands-on/database"
	"github.com/dionomusuko/go-api-hands-on/model"
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
	task, err := database.FindById(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, task)
}

func Create(c echo.Context) error {
	title := c.QueryParam("title")
	description := c.QueryParam("description")
	t := model.Task{
		Title:       title,
		Description: description,
	}
	err := database.Store(t)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.String(http.StatusCreated, "task created")
}

func Delete(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	err := database.Delete(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.String(http.StatusOK, "task deleted")
}

func Update(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	title := c.QueryParam("title")
	description := c.QueryParam("description")
	t := model.Task{
		Id:          id,
		Title:       title,
		Description: description,
	}
	err := database.Update(t)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.String(http.StatusOK, "task updated")
}
