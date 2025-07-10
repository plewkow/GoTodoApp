package main

import (
	"draft-zadania-1/api"
	"draft-zadania-1/config"
	"draft-zadania-1/repo"
	"draft-zadania-1/router"
	"draft-zadania-1/services"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	userRepo := repo.NewUserRepository(config.DB)
	taskRepo := repo.NewTaskRepository(config.DB)
	userService := services.NewUserService(userRepo)
	taskService := services.NewTaskService(taskRepo, userRepo)
	userHandler := &api.UserHandler{Service: userService}
	taskHandler := &api.TaskHandler{Service: taskService}
	e := echo.New()
	router.RegisterRoutes(e, taskHandler, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
	//cli.HandleUserCommands(userService, taskService, os.Args[1:])
}
