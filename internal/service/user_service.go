package service


import (
	"errors"
	"github.com/sabhy92/go-mockery-tutorial/internal/repository"
	"fmt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*repository.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *UserService) RegisterUser(name, email string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if email == "" {
		return errors.New("email cannot be empty")
	}

	user := &repository.User{
		Name: name,
		Email: email,
	}

	return s.repo.CreateUser(user)
}

func (s *UserService) RemoveUser(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	return s.repo.DeleteUser(id)
}