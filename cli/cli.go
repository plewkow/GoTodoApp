package cli

import (
	"draft-zadania-1/models"
	"draft-zadania-1/services"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func HandleUserCommands(userService *services.UserService, taskService *services.TaskService, args []string) {
	switch args[0] {
	case "get_all_users":
		users, err := userService.GetAllUsers()
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		for _, u := range users {
			fmt.Printf("ID: %d\nUsername: %s\nEmail: %s\n\n", u.Id, u.Username, u.Email)
		}
	case "get_user_by_id":
		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("invalid user ID:", err)
			return
		}

		user, err := userService.GetUserById(id)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Printf("ID: %d\nUsername: %s\nEmail: %s\n\n", user.Id, user.Username, user.Email)

	case "delete_user":
		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("invalid user ID:", err)
			return
		}
		err = userService.DeleteUserById(id)
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Printf("user with ID %d deleted\n", id)
		}

	case "create_user":
		file := args[1]
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("error reading file:", err)
			return
		}
		var user models.User
		if err := json.Unmarshal(data, &user); err != nil {
			fmt.Println("error parsing json:", err)
			return
		}
		created, err := userService.CreateUser(user)
		if err != nil {
			fmt.Println("error:", err)
			return
		} else {
			fmt.Printf("user created with ID: %d\n", created.Id)

		}

	case "update_user":
		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("invalid user ID:", err)
			return
		}

		file := args[2]
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("error reading file:", err)
			return
		}

		var updatedUser models.User
		if err := json.Unmarshal(data, &updatedUser); err != nil {
			fmt.Println("invalid JSON:", err)
			return
		}

		updatedUser.Id = id

		updated, err := userService.UpdateUser(updatedUser)
		if err != nil {
			fmt.Println("error:", err)
			return
		} else {
			fmt.Printf("user updated with ID: %d\n", updated.Id)
		}

	case "create_task":
		file := args[1]
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		var task models.Task
		if err := json.Unmarshal(data, &task); err != nil {
			fmt.Println("error:", err)
			return
		}
		created, err := taskService.CreateTask(task)
		if err != nil {
			fmt.Println("error:", err)
			return
		} else {
			fmt.Printf("task created with ID: %d\n", created.Id)
		}

	case "get_all_tasks":
		tasks, err := taskService.GetAllTasks()
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		for _, task := range tasks {
			fmt.Printf("ID: %d, Title: %s, Due Date: %s, UserID: %d, Status: %s\n",
				task.Id, task.Title, task.DueDate.Format("2006-01-02 15:04"), task.UserId, task.Status.String())
		}

	case "get_tasks_by_user":
		userId, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("invalid user ID")
			return
		}
		tasks, err := taskService.GetTasksByUserId(userId)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		for _, task := range tasks {
			fmt.Printf("ID: %d, Title: %s, Due Date: %s, UserID: %d, Status: %s\n",
				task.Id, task.Title, task.DueDate.Format("2006-01-02 15:04"), task.UserId, task.Status.String())
		}

	case "update_task":
		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("invalid task ID:", err)
			return
		}

		file := args[2]
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("error reading file:", err)
			return
		}

		var updatedTask models.Task
		if err := json.Unmarshal(data, &updatedTask); err != nil {
			fmt.Println("invalid JSON:", err)
			return
		}

		updatedTask.Id = id

		updated, err := taskService.UpdateTask(updatedTask)
		if err != nil {
			fmt.Println("error:", err)
			return
		} else {
			fmt.Printf("task updated with ID: %d\n", updated.Id)
		}

	case "delete_task":
		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("invalid task ID:", err)
			return
		}
		err = taskService.DeleteTask(id)
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Printf("task with ID %d deleted\n", id)
		}

	default:
		fmt.Println("unknown command:", args[0])
	}
}
