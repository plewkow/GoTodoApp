package services

import (
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/models"
	"draft-zadania-1/repo"
	"errors"
	"github.com/google/uuid"
)

type UserService struct {
	repo repo.UserRepoInterface
}

func NewUserService(repo repo.UserRepoInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, appErr.ErrInternal
	}
	return createdUser, nil
}

func (s *UserService) UpdateUser(user models.User) (*models.User, error) {
	updatedUser, err := s.repo.Update(user)
	if err != nil {
		if errors.Is(err, appErr.ErrUserNotFound) {
			return nil, appErr.ErrUserNotFound
		}
		return nil, appErr.ErrInternal
	}
	return updatedUser, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return []models.User{}, appErr.ErrInternal
	}
	return users, nil
}

func (s *UserService) GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, appErr.ErrUserNotFound) {
			return nil, appErr.ErrUserNotFound
		}
		return nil, appErr.ErrInternal
	}
	return user, nil
}

func (s *UserService) DeleteUserById(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, appErr.ErrUserNotFound) {
			return appErr.ErrUserNotFound
		}
		return appErr.ErrInternal
	}
	return nil
}
