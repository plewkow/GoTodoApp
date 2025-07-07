package main

import (
	"draft-zadania-1/api"
	"draft-zadania-1/repo"
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

	e.Logger.Fatal(e.Start(":8080"))
	//cli.HandleUserCommands(userService, taskService, os.Args[1:])
}

//func printUsage() {
//	fmt.Println("Usage:")
//	fmt.Println("  create_user <user.json>")
//	fmt.Println("  update_user <user_id> <user.json>")
//	fmt.Println("  delete_user <user_id>")
//	fmt.Println("  get_user_by_id <user_id>")
//	fmt.Println("  get_all_users")
//
//	fmt.Println()
//	fmt.Println("  create_task <task.json>")
//	fmt.Println("  update_task <task_id> <task.json>")
//	fmt.Println("  delete_task <task_id>")
//	fmt.Println("  get_tasks_by_user <user_id>")
//	fmt.Println("  get_all_tasks")
//
//	fmt.Println()
//	fmt.Println("Examples:")
//	fmt.Println("  ./main create_user data/new_user.json")
//	fmt.Println("  ./main update_task 3 data/task_update.json")
//	fmt.Println("  ./main get_tasks_by_user 1")
//	fmt.Println("  ./main delete_user 2")
//
//	fmt.Println()
//	fmt.Println("Notes:")
//	fmt.Println("  - Dates in task JSONs must follow format: \"2006-01-02 15:04\"")
//	fmt.Println("  - JSON files must match the User or Task model structures")
//	fmt.Println("  - To display this help message: ./main --help")
//}
