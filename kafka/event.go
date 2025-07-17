package kafka

import (
	"github.com/google/uuid"
	"time"
)

type UserEvent struct {
	Typ       string    `json:"type"`
	UserId    uuid.UUID `json:"user_id"`
	Timestamp time.Time `json:"due_date"`
}

type TaskEvent struct {
	Typ       string    `json:"type"`
	TaskId    uuid.UUID `json:"task_id"`
	UserId    uuid.UUID `json:"user_id"`
	Timestamp time.Time `json:"due_date"`
}
