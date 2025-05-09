package models

import (
	"time"

	"github.com/askmhs/gin-book-store/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()

	newPassword, _ := helpers.HashPassword(user.Password)

	user.Password = newPassword

	return nil
}

func (user *User) BeforeUpdate(db *gorm.DB) error {
	user.UpdatedAt = time.Now().Local()
	return nil
}
