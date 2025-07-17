package services

import (
	"context"
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/kafka"
	"draft-zadania-1/models"
	"draft-zadania-1/repo"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type TaskService struct {
	repo     *repo.TaskRepository
	userRepo *repo.UserRepository
	kafka    *kafka.KafkaProducer
}

func NewTaskService(repo *repo.TaskRepository, repository *repo.UserRepository, kafka *kafka.KafkaProducer) *TaskService {
	return &TaskService{repo: repo, userRepo: repository, kafka: kafka}
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
	//jsonTask, err := json.Marshal(createdTask)
	//err = s.kafka.Produce(context.Background(), "todo-task", jsonTask)
	event := kafka.TaskEvent{
		Typ:       "CREATE",
		TaskId:    createdTask.Id,
		UserId:    createdTask.UserId,
		Timestamp: time.Now(),
	}

	eventJson, err := json.Marshal(event)

	err = s.kafka.Produce(context.Background(), "todo-task", eventJson)
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
	//jsonTask, err := json.Marshal(updatedTask)
	//err = s.kafka.Produce(context.Background(), "todo-task", jsonTask)
	event := kafka.TaskEvent{
		Typ:       "UPDATE",
		TaskId:    updatedTask.Id,
		UserId:    updatedTask.UserId,
		Timestamp: time.Now(),
	}

	eventJson, err := json.Marshal(event)

	err = s.kafka.Produce(context.Background(), "todo-task", eventJson)
	if err != nil {
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
