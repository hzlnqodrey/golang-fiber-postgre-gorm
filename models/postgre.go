package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host 		string
	Port		string
	Password	string
	User 		string
	DBname		string
	SSLMode		string
}

func NewConnection(config *Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBname, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	// IF Conn success, Create BOOK Table
	db.AutoMigrate(&Books{})

	return db, err 
}