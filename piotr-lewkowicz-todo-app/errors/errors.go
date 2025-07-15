package errors

import "net/http"

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrUserNotFound = &AppError{
		Code:    "USER_NOT_FOUND",
		Message: "User not found",
		Status:  http.StatusNotFound,
	}
	ErrTaskNotFound = &AppError{
		Code:    "TASK_NOT_FOUND",
		Message: "Task not found",
		Status:  http.StatusNotFound,
	}
	ErrInvalidInput = &AppError{
		Code:    "INVALID_INPUT",
		Message: "Invalid input provided",
		Status:  http.StatusBadRequest,
	}
	ErrInternal = &AppError{
		Code:    "INTERNAL_ERROR",
		Message: "Something went wrong",
		Status:  http.StatusInternalServerError,
	}
)
