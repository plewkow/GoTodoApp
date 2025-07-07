package api

import (
	"draft-zadania-1/dto"
	"draft-zadania-1/models"
	"draft-zadania-1/services"
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service *services.UserService
	//Validate *validator.Validate
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var response []dto.ResponseUserDTO
	for _, user := range users {
		response = append(response, dto.ResponseUserDTO{
			Username: user.Username,
			Email:    user.Email,
		})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	user, err := h.Service.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var response []dto.ResponseUserDTO
	response = append(response, dto.ResponseUserDTO{
		Username: user.Username,
		Email:    user.Email,
	})
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var createDto dto.CreateUserDTO
	if err := c.Bind(&createDto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	//if err := h.Validate.Struct(createDto); err != nil {
	//	return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	//}

	user := models.User{
		Username: createDto.Username,
		Email:    createDto.Email,
	}

	created, err := h.Service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	response := dto.ResponseUserDTO{
		Username: created.Username,
		Email:    created.Email,
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	userFromDb, err := h.Service.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var updateDto dto.UpdateUserDTO
	if err := c.Bind(&updateDto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	user := models.User{
		Id:       userFromDb.Id,
		Username: updateDto.Username,
		Email:    updateDto.Email,
	}

	updated, err := h.Service.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	response := dto.ResponseUserDTO{
		Username: updated.Username,
		Email:    updated.Email,
	}

	return c.JSON(http.StatusNoContent, response)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	err = h.Service.DeleteUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := echo.Map{
		"message": "User deleted successfully",
	}

	return c.JSON(http.StatusNoContent, response)
}
