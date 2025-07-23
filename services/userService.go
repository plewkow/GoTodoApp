package services

import (
	"context"
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/kafka"
	"draft-zadania-1/models"
	"draft-zadania-1/repo"
	"draft-zadania-1/utils"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type UserService struct {
	repo  repo.UserRepoInterface
	kafka *kafka.KafkaProducer
}

func NewUserService(repo repo.UserRepoInterface, kafka *kafka.KafkaProducer) *UserService {
	return &UserService{repo: repo, kafka: kafka}
}

func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, appErr.ErrInternal
	}
	//jsonTask, err := json.Marshal(createdUser)

	event := kafka.UserEvent{
		Typ:       "CREATE",
		UserId:    createdUser.Id,
		Timestamp: time.Now(),
	}

	eventJson, err := json.Marshal(event)

	//s.kafka.Enqueue("todo-user", eventJson)

	utils.AddEventToChannel(eventJson)

	err = s.kafka.Produce(context.Background(), "todo-user", eventJson)
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
	//jsonTask, err := json.Marshal(updatedUser)
	//err = s.kafka.Produce(context.Background(), "todo-user", jsonTask)

	event := kafka.UserEvent{
		Typ:       "UPDATE",
		UserId:    updatedUser.Id,
		Timestamp: time.Now(),
	}

	eventJson, err := json.Marshal(event)

	//s.kafka.Enqueue("todo-user", eventJson)
	utils.AddEventToChannel(eventJson)

	err = s.kafka.Produce(context.Background(), "todo-user", eventJson)
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
