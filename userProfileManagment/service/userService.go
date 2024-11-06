package service

import (
	"fmt"
	"userProfileManagment/model"
	"userProfileManagment/repository"
)

// UserService defines methods for business logic related to users
type UserService interface {
	CreateUser(user model.User) error
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User, id uint) error
	DeleteUser(id uint) error
}

// userService is the implementation of UserService
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new instance of userService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser creates a new user
func (s *userService) CreateUser(user model.User) error {
	return s.repo.CreateUser(user)
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(id uint) (*model.User, error) {
	if !s.repo.UserExists(id) {
		return nil, fmt.Errorf("user with ID %d does not exist", id)
	}
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return user, nil
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(user *model.User, id uint) error {
	return s.repo.UpdateUser(user, id)
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(id uint) error {
	if !s.repo.UserExists(id) {
		return fmt.Errorf("user with ID %d does not exist", id)
	}
	err := s.repo.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
