package models

import (
	"gorm.io/gorm"
)

// Books Model
type Books struct {
	ID			uint	`gorm:"primary key; autoIncrement" json:"id"`
	Author		*string	`json:"author"`
	Title		*string `json:"title"`
	Publisher 	*string `json:"publisher"`
}

// Export Book
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})

	return err
}