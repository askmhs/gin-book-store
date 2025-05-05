package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (book *Book) BeforeCreate(db *gorm.DB) error {
	book.CreatedAt = time.Now().Local()
	book.UpdatedAt = time.Now().Local()
	return nil
}

func (book *Book) BeforeUpdate(db *gorm.DB) error {
	book.UpdatedAt = time.Now().Local()
	return nil
}
