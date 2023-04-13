package services

import (
	"fmt"
	"simple-uber/internal/models"
	"simple-uber/internal/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

type SimpleUberUserService interface {
	SaveUser(models.User) (models.User, error)
	FindUserByUserId(userId uint64) (models.User, error)
	DeleteUser(user models.User) error
}

func (s UserService) SaveUser(user models.User) (models.User, error) {
	savedUser, err := s.repository.SaveUser(user)
	if err != nil {
		return models.User{}, fmt.Errorf("user creation failed with error: %w", err)
	}
	return savedUser, nil
}

func (s UserService) FindUserByUserId(userId uint64) (models.User, error) {
	user, err := s.repository.FindUserById(userId)
	if err != nil {
		return models.User{}, fmt.Errorf("account retrieval failed with error: %w", err)
	}
	return user, nil
}

func (s UserService) DeleteUser(user models.User) error {
	err := s.repository.DeleteUser(user)
	if err != nil {
		return fmt.Errorf("user deletion failed with error: %w", err)
	}
	return nil
}

func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}
