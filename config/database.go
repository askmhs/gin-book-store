package config

import (
	"github.com/askmhs/gin-book-store/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(
		&models.Book{},
		&models.User{},
	)

	if err != nil {
		return
	}

	DB = database
}
