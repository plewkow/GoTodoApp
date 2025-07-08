package utils

import (
	"draft-zadania-1/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func WriteAppError(c echo.Context, err error) error {
	if appErr, ok := err.(*errors.AppError); ok {
		return c.JSON(appErr.Status, echo.Map{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
	}
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"code":    "UNKNOWN_ERROR",
		"message": err.Error(),
	})
}
