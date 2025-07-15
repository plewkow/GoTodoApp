package api

import (
	"github.com/labstack/echo/v4"
	openapitypes "github.com/oapi-codegen/runtime/types"
)

type ServerInterface interface {
	ListTasks(ctx echo.Context) error
	CreateTask(ctx echo.Context) error
	GetTasksByUserId(ctx echo.Context, id openapitypes.UUID) error
	DeleteTask(ctx echo.Context, id openapitypes.UUID) error
	GetTaskById(ctx echo.Context, id openapitypes.UUID) error
	UpdateTask(ctx echo.Context, id openapitypes.UUID) error

	ListUsers(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context, id openapitypes.UUID) error
	GetUserById(ctx echo.Context, id openapitypes.UUID) error
	UpdateUser(ctx echo.Context, id openapitypes.UUID) error
}

type combinedHandler struct {
	userHandler *UserHandler
	taskHandler *TaskHandler
}

func NewCombinedHandler(u *UserHandler, t *TaskHandler) ServerInterface {
	return &combinedHandler{
		userHandler: u,
		taskHandler: t,
	}
}

// Metody task'u

func (h *combinedHandler) ListTasks(c echo.Context) error {
	return h.taskHandler.GetAllTasks(c)
}

func (h *combinedHandler) CreateTask(c echo.Context) error {
	return h.taskHandler.CreateTask(c)
}

func (h *combinedHandler) GetTasksByUserId(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.taskHandler.GetTaskByUserId(c)
}

func (h *combinedHandler) DeleteTask(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.taskHandler.DeleteTask(c)
}

func (h *combinedHandler) GetTaskById(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.taskHandler.GetTaskById(c)
}

func (h *combinedHandler) UpdateTask(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.taskHandler.UpdateTask(c)
}

// Metody user'a

func (h *combinedHandler) ListUsers(c echo.Context) error {
	return h.userHandler.GetAllUsers(c)
}

func (h *combinedHandler) CreateUser(c echo.Context) error {
	return h.userHandler.CreateUser(c)
}

func (h *combinedHandler) DeleteUser(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.userHandler.DeleteUser(c)
}

func (h *combinedHandler) GetUserById(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.userHandler.GetUserById(c)
}

func (h *combinedHandler) UpdateUser(c echo.Context, id openapitypes.UUID) error {
	c.SetParamNames("id")
	c.SetParamValues(id.String())
	return h.userHandler.UpdateUser(c)
}
