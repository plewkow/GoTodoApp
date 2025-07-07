package api

import (
	"draft-zadania-1/dto"
	"draft-zadania-1/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	Service *services.TaskService
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := dto.ToResponseTaskDTOs(tasks)
	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var createDto dto.CreateTaskDTO
	if err := c.Bind(&createDto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	//if err := h.Validate.Struct(createDto); err != nil {
	//	return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	//}

	task := dto.ToTaskCreate(createDto)

	created, err := h.Service.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	var response dto.ResponseTaskDTO
	response = dto.ToResponseTaskDTO(created)
	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) GetTaskById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	task, err := h.Service.GetTaskById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//var response []dto.ResponseUserDTO
	//response = append(response, dto.ResponseUserDTO{
	//	Username: user.Username,
	//	Email:    user.Email,
	//})
	var response dto.ResponseTaskDTO
	response = dto.ToResponseTaskDTO(task)
	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) GetTaskByUserId(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	tasks, err := h.Service.GetTasksByUserId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := dto.ToResponseTaskDTOs(tasks)
	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var updateDto dto.UpdateTaskDTO
	if err := c.Bind(&updateDto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	task := dto.ToTask(updateDto)
	task.Id = id

	updated, err := h.Service.UpdateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	response := dto.ToResponseTaskDTO(updated)
	return c.JSON(http.StatusNoContent, response)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	err = h.Service.DeleteTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := echo.Map{
		"message": "Task deleted successfully",
	}

	return c.JSON(http.StatusNoContent, response)
}
