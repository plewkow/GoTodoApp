package api

import (
	"draft-zadania-1/dto"
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/services"
	"draft-zadania-1/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	Service *services.TaskService
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseTaskDTOs(tasks)
	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var createDto dto.TaskDTO
	if err := utils.BindAndValidate(c, &createDto); err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	task := dto.ToTask(createDto)
	created, err := h.Service.CreateTask(task)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseTaskDTO(*created)
	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) GetTaskById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	task, err := h.Service.GetTaskById(id)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseTaskDTO(*task)
	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) GetTaskByUserId(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	tasks, err := h.Service.GetTasksByUserId(id)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseTaskDTOs(tasks)
	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	var updateDto dto.TaskDTO
	if err := utils.BindAndValidate(c, &updateDto); err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	task := dto.ToTask(updateDto)
	task.Id = id
	updated, err := h.Service.UpdateTask(task)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseTaskDTO(*updated)
	return c.JSON(http.StatusNoContent, response)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	err = h.Service.DeleteTask(id)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := echo.Map{
		"message": "Task deleted successfully",
	}
	return c.JSON(http.StatusNoContent, response)
}
