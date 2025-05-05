package interfaces

type BaseRepository[T any] interface {
	FindAll(filters map[string]any) ([]T, error)
	FindById(id uint) (*T, error)
	Create(data *T) error
	Update(id uint, data *T) error
	Delete(id uint) error
}
