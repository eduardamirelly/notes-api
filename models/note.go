package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id          uint
	Title       string
	Description string

	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}
