package api

import (
	"draft-zadania-1/dto"
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/services"
	"draft-zadania-1/utils"

	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseUserDTOs(users)
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	user, err := h.Service.GetUserById(id)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseUserDTO(user)
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var createDto dto.UserDTO
	if err := utils.BindAndValidate(c, &createDto); err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	user := dto.ToUser(createDto)
	created, err := h.Service.CreateUser(user)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseUserDTO(created)
	return c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	var updateDto dto.UserDTO
	if err := utils.BindAndValidate(c, &updateDto); err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	user := dto.ToUser(updateDto)
	user.Id = id
	updated, err := h.Service.UpdateUser(user)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := dto.ToResponseUserDTO(updated)
	return c.JSON(http.StatusNoContent, response)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.WriteAppError(c, appErr.ErrInvalidInput)
	}
	err = h.Service.DeleteUserById(id)
	if err != nil {
		return utils.WriteAppError(c, err)
	}
	response := echo.Map{
		"message": "User deleted successfully",
	}
	return c.JSON(http.StatusNoContent, response)
}
