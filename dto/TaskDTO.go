package dto

import "time"

type CreateTaskDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      Status    `json:"status"`
	UserId      int       `json:"user_id"`
}

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (status Status) String() string {
	return [3]string{"Todo", "InProgress", "Done"}[status]
}

type ResponseTaskDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      Status    `json:"status"`
	UserId      int       `json:"user_id"`
}

type UpdateTaskDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      int       `json:"status"`
	UserId      int       `json:"user_id"`
}
