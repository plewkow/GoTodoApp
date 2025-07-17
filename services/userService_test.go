package services

import (
	"draft-zadania-1/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (repo *MockUserRepository) GetAll() ([]models.User, error) {
	args := repo.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (repo *MockUserRepository) GetById(id uuid.UUID) (*models.User, error) {
	args := repo.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (repo *MockUserRepository) Create(user models.User) (*models.User, error) {
	args := repo.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (repo *MockUserRepository) Update(user models.User) (*models.User, error) {
	args := repo.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (repo *MockUserRepository) Delete(id uuid.UUID) error {
	args := repo.Called(id)
	return args.Error(0)
}

func TestCreateUserSuccess(t *testing.T) {
	mockUserRepository := new(MockUserRepository)
	userService := NewUserService(mockUserRepository)

	user := models.User{
		Id:       uuid.New(),
		Username: "Jazzman",
		Email:    "jazzman@email.com",
	}

	mockUserRepository.On("Create", user).Return(&user, nil)

	createdUser, err := userService.CreateUser(user)

	assert.NoError(t, err)

	assert.Equal(t, user, *createdUser)

	assert.Equal(t, user.Id, createdUser.Id)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Email, createdUser.Email)

	mockUserRepository.AssertExpectations(t)
}
