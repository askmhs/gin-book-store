package services

import (
	"errors"

	"github.com/askmhs/gin-book-store/helpers"
	"github.com/askmhs/gin-book-store/models"
	"github.com/askmhs/gin-book-store/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
	jwtService *JwtService
}

func NewUserService(userRepo *repositories.UserRepository, jwtService *JwtService) *UserService {
	return &UserService{
		repository: userRepo,
		jwtService: jwtService,
	}
}

func (s *UserService) CreateUser(data *models.User) error {
	return s.repository.Create(data)
}

func (s *UserService) LoginUser(username, password string) (any, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if passwordMatch := helpers.VerifyPassword(password, user.Password); !passwordMatch {
		return nil, errors.New("username and password combination doesn't match")
	}

	userData := map[string]any{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"userName":  user.UserName,
	}

	token, err := s.jwtService.GenerateToken(userData)
	if err != nil {
		return nil, err
	}

	return token, nil
}
