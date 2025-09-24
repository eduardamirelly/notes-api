package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint
	Username string
	Email    string

	CreatedAt time.Time
	UpdatedAt time.Time

	Notes []Note `gorm:"foreignKey:UserID"`
}
