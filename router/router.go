package router

import (
	"github.com/dionomusuko/go-api-hands-on/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")
	task := v1.Group("/tasks")
	task.GET("", controller.ListTasks)
	task.GET("/:id", controller.FindById)
	task.POST("", controller.Create)
	task.PUT("/:id", controller.Update)
	task.DELETE("/:id", controller.Delete)
	return e
}
