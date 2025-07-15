package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
