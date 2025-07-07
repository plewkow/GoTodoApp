package repo

import (
	"draft-zadania-1/models"
	"encoding/json"
	"fmt"
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
		return nil, err
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
		return err
	}
	if err := json.Unmarshal(data, &r.tasks); err != nil {
		return err
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
	data, err := json.Marshal(r.tasks)
	if err != nil {
		return err
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
	return models.Task{}, fmt.Errorf("task not found")
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
		return fmt.Errorf("task with id %d not found", id)
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
		}
	}
	return result, fmt.Errorf("task not found")
}

func (r *TaskRepository) GetByUserId(userId int) ([]models.Task, error) {
	var result []models.Task
	for _, t := range r.tasks {
		if t.UserId == userId {
			result = append(result, t)
		}
	}
	return result, fmt.Errorf("user not found")
}
