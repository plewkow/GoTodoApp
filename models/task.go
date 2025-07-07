package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      Status    `json:"status"`
	UserId      int       `json:"user_id"`
}
