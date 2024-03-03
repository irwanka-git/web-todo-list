package repository

import (
	"errors"
	"irwanka/webtodolist/entity"
)

type TaskRepository interface {
	GetListTask(user_id int32) ([]*entity.Task, error)
	GetDetilTask(id_task int32) (*entity.Task, error)
	CreateTask(task entity.Task) (*entity.Task, error)
	UpdateTask(task entity.Task) error
	DeleteTask(id_task int32) error
}

func NewTaskRepository() TaskRepository {
	return &repo{}
}

func (*repo) GetListTask(user_id int32) ([]*entity.Task, error) {
	var data []*entity.Task
	result := db.Table("task").Where("create_by = ? ", user_id).Order("created_at desc").Scan(&data)
	if result.RowsAffected == 0 {
		return []*entity.Task{}, errors.New("task not found")
	}
	return data, nil
}

func (*repo) GetDetilTask(id_task int32) (*entity.Task, error) {
	var data *entity.Task
	result := db.Table("task").Where("id_task = ? ", id_task).First(&data)
	if result.RowsAffected == 0 {
		return nil, errors.New("task not found")
	}
	return data, nil
}

func (*repo) CreateTask(task entity.Task) (*entity.Task, error) {
	tx := db.Begin()
	tx.Table("task").Create(&task)
	err := tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (*repo) UpdateTask(task entity.Task) error {
	tx := db.Begin()
	tx.Table("task").Where("id_task = ?", &task.IDTask).Updates(&task)
	return tx.Commit().Error
}

func (*repo) DeleteTask(id_task int32) error {
	var task *entity.Task
	tx := db.Begin()
	tx.Table("task").Where("id_task = ?", id_task).Delete(&task)
	return tx.Commit().Error
}
