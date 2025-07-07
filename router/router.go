package router

import (
	"draft-zadania-1/api"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, taskHandler *api.TaskHandler, userHandler *api.UserHandler) {
	e.GET("/api/tasks", taskHandler.GetAllTasks)
	e.GET("/api/tasks/:id", taskHandler.GetTaskById)
	e.GET("/api/tasks/user/task", taskHandler.GetTaskById)
	e.POST("/api/tasks", taskHandler.CreateTask)
	e.PUT("/api/tasks/:id", taskHandler.UpdateTask)
	e.DELETE("/api/tasks/:id", taskHandler.DeleteTask)
	e.GET("/api/users", userHandler.GetAllUsers)
	e.PUT("/api/users/:id", userHandler.UpdateUser)
	e.GET("/api/users/:id", userHandler.GetUserById)
	e.POST("/api/users", userHandler.CreateUser)
	e.DELETE("/api/users/:id", userHandler.DeleteUser)
}
