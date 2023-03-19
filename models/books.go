package models

import (
	"gorm.io/gorm"
)

// Books Model
type Books struct {
	ID          uint    `gorm:"primaryKey; autoIncrement" json:"id"`
	Title       *string `gorm:"type:varchar(300)" json:"title"`
	Description *string `gorm:"type:varchar(500)" json:"description"`
	Author      *string `gorm:"type:varchar(300)" json:"author"`
	Publisher   *string `gorm:"type:varchar(300)" json:"publisher"`
}

// Export Book
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})

	return err
}
