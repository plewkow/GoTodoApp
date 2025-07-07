package main

import (
	"fmt"
	"os"

	"draft-zadania-1/cli"
	"draft-zadania-1/repo"
	"draft-zadania-1/services"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		printUsage()
		return
	}

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
	cli.HandleUserCommands(userService, taskService, os.Args[1:])
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  create_user <user.json>")
	fmt.Println("  update_user <user_id> <user.json>")
	fmt.Println("  delete_user <user_id>")
	fmt.Println("  get_user_by_id <user_id>")
	fmt.Println("  get_all_users")

	fmt.Println()
	fmt.Println("  create_task <task.json>")
	fmt.Println("  update_task <task_id> <task.json>")
	fmt.Println("  delete_task <task_id>")
	fmt.Println("  get_tasks_by_user <user_id>")
	fmt.Println("  get_all_tasks")

	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ./main create_user data/new_user.json")
	fmt.Println("  ./main update_task 3 data/task_update.json")
	fmt.Println("  ./main get_tasks_by_user 1")
	fmt.Println("  ./main delete_user 2")

	fmt.Println()
	fmt.Println("Notes:")
	fmt.Println("  - Dates in task JSONs must follow format: \"2006-01-02 15:04\"")
	fmt.Println("  - JSON files must match the User or Task model structures")
	fmt.Println("  - To display this help message: ./main --help")
}
