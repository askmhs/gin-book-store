package services

import (
	"github.com/askmhs/gin-book-store/interfaces"
	"github.com/askmhs/gin-book-store/models"
)

type BookService struct {
	repository interfaces.BaseRepository[models.Book]
}

func NewBookService(repo interfaces.BaseRepository[models.Book]) *BookService {
	return &BookService{repository: repo}
}

func (bs *BookService) GetBooks(filter map[string]any) ([]models.Book, error) {
	return bs.repository.FindAll(filter)
}

func (bs *BookService) GetBook(id uint) (*models.Book, error) {
	return bs.repository.FindById(id)
}

func (bs *BookService) CreateBook(data *models.Book) error {
	return bs.repository.Create(data)
}

func (bs *BookService) UpdateBook(id uint, data *models.Book) error {
	return bs.repository.Update(id, data)
}

func (bs *BookService) DeleteBook(id uint) error {
	return bs.repository.Delete(id)
}
