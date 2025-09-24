package main

import (
	"os"
	"path/filepath"

	"github.com/eduardamirelly/notes-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	currentDir, err := os.Getwd()

	if err != nil {
		panic("failed to get current directory")
	}

	dbPath := filepath.Join(currentDir, "notes_api.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Note{})
}
