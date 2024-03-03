package service

import (
	"irwanka/webtodolist/entity"
	"irwanka/webtodolist/repository"
)

type TaskService interface {
	GetListTask(user_id int32) ([]*entity.Task, error)
	GetDetilTask(id_task int32) (*entity.Task, error)
	CreateTask(task entity.Task) (*entity.Task, error)
	UpdateTask(task entity.Task) error
	DeleteTask(id_task int32) error
}

var (
	taskRepo repository.TaskRepository
)

func NewTaskService(repo repository.TaskRepository) TaskService {
	taskRepo = repo
	return &service{}
}

func (*service) GetListTask(user_id int32) ([]*entity.Task, error) {
	return taskRepo.GetListTask(user_id)
}

func (*service) GetDetilTask(id_task int32) (*entity.Task, error) {
	return taskRepo.GetDetilTask(id_task)
}

func (*service) CreateTask(task entity.Task) (*entity.Task, error) {
	return taskRepo.CreateTask(task)
}

func (*service) UpdateTask(task entity.Task) error {
	return taskRepo.UpdateTask(task)
}

func (*service) DeleteTask(id_task int32) error {
	return taskRepo.DeleteTask(id_task)
}
