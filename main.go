package main

import (
	"draft-zadania-1/api"
	"draft-zadania-1/config"
	"draft-zadania-1/repo"
	spec "draft-zadania-1/spec"
	//"draft-zadania-1/router"
	"draft-zadania-1/services"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config.InitDB()
	userRepo := repo.NewUserRepository(config.DB)
	taskRepo := repo.NewTaskRepository(config.DB)
	userService := services.NewUserService(userRepo)
	taskService := services.NewTaskService(taskRepo, userRepo)
	userHandler := &api.UserHandler{Service: userService}
	taskHandler := &api.TaskHandler{Service: taskService}
	combined := api.NewCombinedHandler(userHandler, taskHandler)
	e := echo.New()
	apiGroup := e.Group("/api")
	e.Static("/api", "dist")
	e.Static("/swagger-ui", "swagger-ui")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	spec.RegisterHandlers(apiGroup, combined)
	//router.RegisterRoutes(e, taskHandler, userHandler) // wcześniej gdy nie było openapi
	e.Logger.Fatal(e.Start(":8080"))
	//cli.HandleUserCommands(userService, taskService, os.Args[1:]) // wcześniej do cli
}
