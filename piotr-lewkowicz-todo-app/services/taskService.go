package services

import (
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/models"
	"draft-zadania-1/repo"
	"errors"
	"github.com/google/uuid"
)

type TaskService struct {
	repo     *repo.TaskRepository
	userRepo *repo.UserRepository
}

func NewTaskService(repo *repo.TaskRepository, repository *repo.UserRepository) *TaskService {
	return &TaskService{repo: repo, userRepo: repository}
}

func (s *TaskService) CreateTask(task models.Task) (*models.Task, error) {
	_, err := s.userRepo.GetById(task.UserId)
	if err != nil {
		if errors.Is(err, appErr.ErrUserNotFound) {
			return nil, err
		}
		return nil, appErr.ErrInternal
	}

	createdTask, err := s.repo.Create(task) // :=
	if err != nil {
		return nil, appErr.ErrInternal
	}
	return createdTask, nil
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return []models.Task{}, appErr.ErrInternal
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(task models.Task) (*models.Task, error) {
	updatedTask, err := s.repo.Update(task)
	if err != nil {
		if errors.Is(err, appErr.ErrTaskNotFound) {
			return nil, err
		}
		return nil, appErr.ErrInternal
	}
	return updatedTask, nil
}

func (s *TaskService) DeleteTask(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, appErr.ErrTaskNotFound) {
			return err
		}
		return appErr.ErrInternal
	}
	return nil
}

func (s *TaskService) GetTasksByUserId(userId uuid.UUID) ([]models.Task, error) {
	tasks, err := s.repo.GetByUserId(userId)
	if err != nil {
		if errors.Is(err, appErr.ErrTaskNotFound) {
			return nil, err
		}
		return nil, appErr.ErrInternal
	}
	return tasks, nil
}

func (s *TaskService) GetTaskById(id uuid.UUID) (*models.Task, error) {
	task, err := s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, appErr.ErrTaskNotFound) {
			return nil, appErr.ErrTaskNotFound
		}
		return nil, appErr.ErrInternal
	}
	return task, nil
}
