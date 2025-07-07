package services

import (
	"draft-zadania-1/models"
	"draft-zadania-1/repo"
)

type TaskService struct {
	repo *repo.TaskRepository
}

func NewTaskService(repo *repo.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	return s.repo.Create(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) UpdateTask(task models.Task) (models.Task, error) {
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}

func (s *TaskService) GetTasksByUserId(userId int) ([]models.Task, error) {
	return s.repo.GetByUserId(userId)
}

func (s *TaskService) GetTaskById(id int) (models.Task, error) {
	return s.repo.GetById(id)
}
