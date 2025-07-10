package repo

import (
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, appErr.ErrInternal
	}
	return &user, nil
}

func (r *UserRepository) Update(user models.User) (*models.User, error) {
	var existingUser models.User
	if err := r.db.First(&existingUser, user.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrUserNotFound
		}
		return nil, appErr.ErrInternal
	}

	if err := r.db.Save(&user).Error; err != nil {
		return nil, appErr.ErrInternal
	}
	return &user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, appErr.ErrInternal
	}
	return users, nil
}

func (r *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrUserNotFound
		}
		return nil, appErr.ErrInternal
	}
	return &user, nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErr.ErrUserNotFound
		}
		return appErr.ErrInternal
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return appErr.ErrInternal
	}
	return nil
}
