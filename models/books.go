package models

import (
	"gorm.io/gorm"
)

// Books Model
type Books struct {
	ID          uint    `gorm:"primaryKey; autoIncrement" json:"id"`
	Title       *string `gorm:"type:varchar(300)" json:"title"`
	Description *string `gorm:"type:text" json:"description"`
	Author      *string `gorm:"type:varchar(300)" json:"author"`
	PublishDate *string `gorm:"type:date" json:"publish_date"`
}

// Export Book
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})

	return err
}
