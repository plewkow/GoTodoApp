package router

//
//import (
//	"draft-zadania-1/api"
//	generated "draft-zadania-1/spec/generated" // wygenerowany pakiet z oapi-codegen
//	"github.com/labstack/echo/v4"
//)
//
//func RegisterRoutes(e *echo.Echo, taskHandler *api.TaskHandler, userHandler *api.UserHandler) {
//	// TASKS
//	//e.GET("/api/tasks", taskHandler.GetAllTasks)
//	//e.GET("/api/tasks/:id", taskHandler.GetTaskById)
//	//e.GET("/api/tasks/user/:id", taskHandler.GetTaskByUserId)
//	//e.POST("/api/tasks", taskHandler.CreateTask)
//	//e.PUT("/api/tasks/:id", taskHandler.UpdateTask)
//	//e.DELETE("/api/tasks/:id", taskHandler.DeleteTask)
//	//// USERS
//	//e.GET("/api/users", userHandler.GetAllUsers)
//	//e.PUT("/api/users/:id", userHandler.UpdateUser)
//	//e.GET("/api/users/:id", userHandler.GetUserById)
//	//e.POST("/api/users", userHandler.CreateUser)
//	//e.DELETE("/api/users/:id", userHandler.DeleteUser)
//
//	type combinedHandler struct {
//		*api.UserHandler
//		*api.TaskHandler
//	}
//
//	handler := &combinedHandler{
//		UserHandler: userHandler,
//		TaskHandler: taskHandler,
//	}
//
//	generated.RegisterHandlers(e, handler)
//}
