package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      Status    `json:"status"`
	UserId      uuid.UUID `gorm:"not null" json:"user_id"`
}
