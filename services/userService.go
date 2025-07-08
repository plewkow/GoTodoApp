package services

import (
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/models"
	"draft-zadania-1/repo"
	"errors"
)

type UserService struct {
	repo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	user, err := s.repo.Create(user)
	if err != nil {
		return models.User{}, appErr.ErrInternal
	}
	return user, nil
}

func (s *UserService) UpdateUser(user models.User) (models.User, error) {
	user, err := s.repo.Update(user)
	if err != nil {
		if errors.Is(err, appErr.ErrUserNotFound) {
			return models.User{}, appErr.ErrUserNotFound
		}
		return models.User{}, appErr.ErrInternal
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return []models.User{}, appErr.ErrInternal
	}
	return users, nil
}

func (s *UserService) GetUserById(id int) (models.User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, appErr.ErrUserNotFound) {
			return models.User{}, appErr.ErrUserNotFound
		}
		return models.User{}, appErr.ErrInternal
	}
	return user, nil
}

func (s *UserService) DeleteUserById(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, appErr.ErrUserNotFound) {
			return appErr.ErrUserNotFound
		}
		return appErr.ErrInternal
	}
	return nil
}
