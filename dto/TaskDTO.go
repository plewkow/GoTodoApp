package dto

import "time"

type TaskDTO struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	Status      int       `json:"status" validate:"required"`
	UserId      int       `json:"user_id" validate:"required"`
}

type ResponseTaskDTO struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
	UserId      int       `json:"user_id"`
}
