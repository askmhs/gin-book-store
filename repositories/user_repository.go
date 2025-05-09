package repositories

import (
	"github.com/askmhs/gin-book-store/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	err := r.DB.Where("user_name = ?", username).First(&user).Error

	return &user, err
}

func (r *UserRepository) Create(data *models.User) error {
	return r.DB.Create(data).Error
}
