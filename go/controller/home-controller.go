package controller

import (
	"encoding/json"
	"irwanka/webtodolist/helper"
	"net/http"
)

type HomeController interface {
	Welcome(w http.ResponseWriter, r *http.Request)
}

func NewHomeController() HomeController {
	return &controller{}
}

func (*controller) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.ResponseMessage{Status: true, Message: "Selamat Datang di Api Web Todolist"})
}
