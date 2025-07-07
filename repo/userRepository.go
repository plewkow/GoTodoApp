package repo

import (
	"draft-zadania-1/models"
	"encoding/json"
	"fmt"
	"os"
)

type UserRepository struct {
	users  []models.User
	nextId int
	file   string
}

func NewUserRepository() (*UserRepository, error) {
	r := &UserRepository{
		file:   "data/users.json",
		nextId: 0,
		users:  []models.User{},
	}
	if err := r.load(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *UserRepository) save() error {
	data, err := json.Marshal(r.users)
	if err != nil {
		return err
	}
	return os.WriteFile(r.file, data, 0644)
}

func (r *UserRepository) load() error {
	data, err := os.ReadFile(r.file)
	if err != nil {
		if os.IsNotExist(err) {
			r.users = []models.User{}
			return r.save()
		}
		return err
	}
	if err := json.Unmarshal(data, &r.users); err != nil {
		return err
	}
	maxId := 0
	for _, user := range r.users {
		if user.Id > maxId {
			maxId = user.Id
		}
	}
	r.nextId = maxId + 1
	return nil
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	user.Id = r.nextId
	r.nextId++
	r.users = append(r.users, user)
	err := r.save()
	return user, err
}

func (r *UserRepository) Update(user models.User) (models.User, error) {
	for i, u := range r.users {
		if u.Id == user.Id {
			r.users[i] = user
			err := r.save()
			return user, err
		}
	}
	return models.User{}, fmt.Errorf("user not found")
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	return r.users, nil
}

func (r *UserRepository) GetById(id int) (models.User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return user, nil
		}
	}
	return models.User{}, fmt.Errorf("user with ID %d not found", id)
}

func (r *UserRepository) Delete(id int) error {
	index := -1
	for i, user := range r.users {
		if user.Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("user with id %d not found", id)
	}
	r.users = append(r.users[:index], r.users[index+1:]...)
	return r.save()
}
