package repo

import (
	appErr "draft-zadania-1/errors"
	"draft-zadania-1/models"
	"encoding/json"
	"os"
)

type TaskRepository struct {
	tasks  []models.Task
	nextId int
	file   string
}

func NewTaskRepository() (*TaskRepository, error) {
	r := &TaskRepository{
		file:   "data/tasks.json",
		tasks:  []models.Task{},
		nextId: 1,
	}
	if err := r.load(); err != nil {
		return nil, appErr.ErrInternal
	}
	return r, nil
}

func (r *TaskRepository) load() error {
	data, err := os.ReadFile(r.file)
	if err != nil {
		if os.IsNotExist(err) {
			r.tasks = []models.Task{}
			return r.save()
		}
		return appErr.ErrInternal
	}
	if err := json.Unmarshal(data, &r.tasks); err != nil {
		return appErr.ErrInternal
	}
	maxId := 0
	for _, task := range r.tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	r.nextId = maxId + 1
	return nil
}

func (r *TaskRepository) save() error {
	data, err := json.MarshalIndent(r.tasks, "", " ")
	if err != nil {
		return appErr.ErrInternal
	}
	return os.WriteFile(r.file, data, 0644)
}

func (r *TaskRepository) Create(task models.Task) (models.Task, error) {
	task.Id = r.nextId
	r.nextId++
	r.tasks = append(r.tasks, task)
	err := r.save()
	return task, err
}

func (r *TaskRepository) Update(task models.Task) (models.Task, error) {
	for i, t := range r.tasks {
		if t.Id == task.Id {
			r.tasks[i] = task
			err := r.save()
			return task, err
		}
	}
	return models.Task{}, appErr.ErrTaskNotFound
}

func (r *TaskRepository) Delete(id int) error {
	index := -1
	for i, task := range r.tasks {
		if task.Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return appErr.ErrTaskNotFound
	}
	r.tasks = append(r.tasks[:index], r.tasks[index+1:]...)
	return r.save()
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	return r.tasks, nil
}

func (r *TaskRepository) GetById(id int) (models.Task, error) {
	var result models.Task
	for _, task := range r.tasks {
		if task.Id == id {
			result = task
			return result, nil
		}
	}
	return models.Task{}, appErr.ErrTaskNotFound
}

func (r *TaskRepository) GetByUserId(userId int) ([]models.Task, error) {
	var result []models.Task
	for _, t := range r.tasks {
		if t.UserId == userId {
			result = append(result, t)
			return result, nil
		}
	}
	return result, appErr.ErrTaskNotFound
}
