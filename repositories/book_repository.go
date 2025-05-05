package repositories

import (
	"fmt"

	"github.com/askmhs/gin-book-store/helpers"
	"github.com/askmhs/gin-book-store/interfaces"
	"github.com/askmhs/gin-book-store/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) interfaces.BaseRepository[models.Book] {
	return &BookRepository{DB: db}
}

func (r *BookRepository) GetFilterableFileds() []string {
	return []string{"title", "author"}
}

func (r *BookRepository) FindAll(filters map[string]any) ([]models.Book, error) {
	var books []models.Book
	query := r.DB.Model(&models.Book{})

	fileterableFields := r.GetFilterableFileds()

	// Apply filters if any
	for key, value := range filters {
		if helpers.Contains(fileterableFields, key) {
			strVal := fmt.Sprintf("%v", value)
			query = query.Where(fmt.Sprintf("%s LIKE ?", key), "%"+strVal+"%")
		}
	}

	err := query.Find(&books).Error
	return books, err
}

func (r *BookRepository) FindById(id uint) (*models.Book, error) {
	var book models.Book

	err := r.DB.First(&book, id).Error

	return &book, err
}

func (r *BookRepository) Create(data *models.Book) error {
	return r.DB.Create(data).Error
}

func (r *BookRepository) Update(id uint, data *models.Book) error {
	return r.DB.Model(&models.Book{}).Where("id = ?", id).Updates(data).Error
}

func (r *BookRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Book{}, id).Error
}
