package controller

import (
	"encoding/json"
	"irwanka/webtodolist/entity"
	"irwanka/webtodolist/helper"
	"irwanka/webtodolist/middleware"
	"irwanka/webtodolist/repository"
	"irwanka/webtodolist/service"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

var (
	taskRepository repository.TaskRepository = repository.NewTaskRepository()
	taskService    service.TaskService       = service.NewTaskService(taskRepository)
)

type TaskController interface {
	GetListTask(w http.ResponseWriter, r *http.Request)
	GetDetilTask(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

func NewTaskController() TaskController {
	return &controller{}
}

func (*controller) GetListTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	ctx := r.Context()
	user_id_context, _ := ctx.Value(helper.VALUE_CONTEXT).(middleware.ValueContext).GetValueContext(helper.USER_ID).(string)
	user_id, _ := strconv.Atoi(user_id_context)
	listTask, err := taskService.GetListTask(int32(user_id))
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Data: listTask})
}

func (*controller) GetDetilTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	ctx := r.Context()
	user_id_context, _ := ctx.Value(helper.VALUE_CONTEXT).(middleware.ValueContext).GetValueContext(helper.USER_ID).(string)
	user_id, _ := strconv.Atoi(user_id_context)

	task_id, err_id := strconv.Atoi(chi.URLParam(r, "id"))
	if err_id != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Request ID Tidak Valid"})
		return
	}
	detil, err_task := taskService.GetDetilTask(int32(task_id))
	if err_task != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: err_task.Error()})
		return
	}
	if detil.CreateBy != int32(user_id) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Akses Tidak Valid"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Data: detil})
}

func (*controller) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	ctx := r.Context()
	user_id_context, _ := ctx.Value(helper.VALUE_CONTEXT).(middleware.ValueContext).GetValueContext(helper.USER_ID).(string)
	user_id, _ := strconv.Atoi(user_id_context)

	var inputTask entity.InputTask
	err := json.NewDecoder(r.Body).Decode(&inputTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "data task wajib diisi", Status: false})
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	errInput := validate.Struct(inputTask)
	if errInput != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errInput.Error(), Status: false})
		return
	}

	var task entity.Task
	task.Title = inputTask.Title
	task.Description = inputTask.Description
	task.CreateBy = int32(user_id)
	task.CreatedAt = time.Now()

	result, errCreate := taskService.CreateTask(task)
	if errCreate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errCreate.Error()})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Message: "task berhasil ditambahkan", Data: result})
}

func (*controller) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	ctx := r.Context()
	user_id_context, _ := ctx.Value(helper.VALUE_CONTEXT).(middleware.ValueContext).GetValueContext(helper.USER_ID).(string)
	user_id, _ := strconv.Atoi(user_id_context)

	task_id, err_id := strconv.Atoi(chi.URLParam(r, "id"))
	if err_id != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Request Tidak Valid"})
		return
	}

	var inputTask entity.InputTask
	err := json.NewDecoder(r.Body).Decode(&inputTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "data task wajib diisi", Status: false})
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	errInput := validate.Struct(inputTask)
	if errInput != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errInput.Error(), Status: false})
		return
	}

	//cek current task
	currentTask, errCek := taskService.GetDetilTask(int32(task_id))
	if errCek != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errCek.Error(), Status: false})
		return
	}
	//cek kepemilikan task
	if currentTask.CreateBy != int32(user_id) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "akses tidak valid", Status: false})
		return
	}

	var updateTask entity.Task
	updateTask.IDTask = currentTask.IDTask
	updateTask.CreateBy = currentTask.CreateBy
	updateTask.CreatedAt = currentTask.CreatedAt
	updateTask.UpdatedAt = time.Now()

	updateTask.Title = inputTask.Title
	updateTask.Description = inputTask.Description

	errUpdate := taskService.UpdateTask(updateTask)
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errUpdate.Error()})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(helper.ResponseData{Status: true, Message: "task berhasil diperbarui", Data: updateTask})
}

func (*controller) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	ctx := r.Context()
	user_id_context, _ := ctx.Value(helper.VALUE_CONTEXT).(middleware.ValueContext).GetValueContext(helper.USER_ID).(string)
	user_id, _ := strconv.Atoi(user_id_context)

	task_id, err_id := strconv.Atoi(chi.URLParam(r, "id"))
	if err_id != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Request Tidak Valid"})
		return
	}

	//cek current task
	currentTask, errCek := taskService.GetDetilTask(int32(task_id))
	if errCek != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errCek.Error(), Status: false})
		return
	}
	//cek kepemilikan task
	if currentTask.CreateBy != int32(user_id) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "akses tidak valid", Status: false})
		return
	}

	errDelete := taskService.DeleteTask(currentTask.IDTask)
	if errDelete != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errDelete.Error()})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(helper.ResponseMessage{Status: true, Message: "task berhasil dihapus"})
}
