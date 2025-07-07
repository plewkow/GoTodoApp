package services

import (
	"draft-zadania-1/models"
	"draft-zadania-1/repo"
)

type UserService struct {
	repo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(user models.User) (models.User, error) {
	return s.repo.Update(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUserById(id int) (models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) DeleteUserById(id int) error {
	return s.repo.Delete(id)
}
