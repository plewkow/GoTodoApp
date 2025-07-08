package services

import (
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/models"
	"draft-zadania-1/repo"
	"errors"
)

type TaskService struct {
	repo *repo.TaskRepository
}

func NewTaskService(repo *repo.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	task, err := s.repo.Create(task)
	if err != nil {
		return models.Task{}, appErr.ErrInternal
	}
	return task, nil
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return []models.Task{}, appErr.ErrInternal
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(task models.Task) (models.Task, error) {
	task, err := s.repo.Update(task)
	if err != nil {
		if errors.Is(err, appErr.ErrTaskNotFound) {
			return models.Task{}, err
		}
		return models.Task{}, appErr.ErrInternal
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, appErr.ErrTaskNotFound) {
			return err
		}
		return appErr.ErrInternal
	}
	return nil
}

func (s *TaskService) GetTasksByUserId(userId int) ([]models.Task, error) {
	task, err := s.repo.GetByUserId(userId)
	if err != nil {
		if errors.Is(err, appErr.ErrTaskNotFound) {
			return nil, err
		}
		return nil, appErr.ErrInternal
	}
	return task, nil
}

func (s *TaskService) GetTaskById(id int) (models.Task, error) {
	task, err := s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, appErr.ErrTaskNotFound) {
			return models.Task{}, appErr.ErrTaskNotFound
		}
		return models.Task{}, appErr.ErrInternal
	}
	return task, nil
}
