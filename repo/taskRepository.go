package repo

import (
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task models.Task) (*models.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, appErr.ErrInternal
	}
	return &task, nil
}

func (r *TaskRepository) Update(task models.Task) (*models.Task, error) {
	var existingTask models.Task
	if err := r.db.First(&existingTask, task.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrTaskNotFound
		}
		return nil, appErr.ErrInternal
	}

	if err := r.db.Save(&task).Error; err != nil {
		return nil, appErr.ErrInternal
	}
	return &task, nil
}

func (r *TaskRepository) Delete(id uuid.UUID) error {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErr.ErrTaskNotFound
		}
		return appErr.ErrInternal
	}

	if err := r.db.Delete(&task).Error; err != nil {
		return appErr.ErrInternal
	}
	return nil
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, appErr.ErrInternal
	}
	return tasks, nil
}

func (r *TaskRepository) GetById(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	if err := r.db.Where("id = ?", id).First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrTaskNotFound
		}
		return nil, appErr.ErrInternal
	}
	return &task, nil
}

func (r *TaskRepository) GetByUserId(userId uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		return nil, appErr.ErrInternal
	}
	if len(tasks) == 0 {
		return nil, appErr.ErrTaskNotFound
	}
	return tasks, nil
}
