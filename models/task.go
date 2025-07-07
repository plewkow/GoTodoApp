package models

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	DueDate     time.Time
	Status      Status
	UserId      int
}
