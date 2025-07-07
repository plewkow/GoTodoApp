package main

import (
	"draft-zadania-1/api"
	"draft-zadania-1/repo"
	"draft-zadania-1/router"
	"draft-zadania-1/services"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	userRepo, err := repo.NewUserRepository()
	if err != nil {
		fmt.Println("failed to initialize repo:", err)
		return
	}
	taskRepo, err := repo.NewTaskRepository()
	if err != nil {
		fmt.Println("failed to initialize repo:", err)
		return
	}

	userService := services.NewUserService(userRepo)
	taskService := services.NewTaskService(taskRepo)

	userHandler := &api.UserHandler{Service: userService}
	taskHandler := &api.TaskHandler{Service: taskService}

	e := echo.New()

	router.RegisterRoutes(e, taskHandler, userHandler)

	e.Logger.Fatal(e.Start(":8080"))
	//cli.HandleUserCommands(userService, taskService, os.Args[1:])
}
