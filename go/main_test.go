package main_test

import (
	"bytes"
	"irwanka/webtodolist/controller"
	"irwanka/webtodolist/helper"
	customMiddleware "irwanka/webtodolist/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/assert"
)

var (
	tokenAuth      *jwtauth.JWTAuth
	homeController controller.HomeController       = controller.NewHomeController()
	authController controller.AuthController       = controller.NewAuthController()
	taskController controller.TaskController       = controller.NewTaskController()
	userMiddleware customMiddleware.UserMiddleware = customMiddleware.NewUserMiddleware()
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/login", authController.SubmitLogin)
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Use(userMiddleware.SetValueContext)
			r.Get("/list-task", taskController.GetListTask)
			r.Get("/get-detil-task/{id}", taskController.GetDetilTask)
			r.Post("/create-task", taskController.CreateTask)
			r.Patch("/update-task/{id}", taskController.UpdateTask)
			r.Delete("/delete-task/{id}", taskController.DeleteTask)
		})
	})
	return r
}

func TestLogin(t *testing.T) {
	t.Logf("Test Login with valid credential")
	testUser := `{"email":"irwanka.email@gmail.com" , "password":"123456"}`
	request, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(testUser))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "OK response is expected")
}
func TestInvalidLogin(t *testing.T) {
	t.Logf("Test Login with invalid credential")
	testUser := `{"email":"irwanka.email@gmail.com" , "password":"xxxxxx"}`
	request, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(testUser))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
}

func TestCreateTask(t *testing.T) {
	t.Logf("Test Create Task")
	//user login untuk mendapatkan JWT Token user

	testUser := `{"email":"irwanka.email@gmail.com" , "password":"123456"}`
	request, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(testUser))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "OK response is expected")

	jwtToken, _ := helper.CreateJWTTokenLogin(userLogin)
	var bearer = "Bearer " + jwtToken

	testTask := `{"title":"Testing Task" , "description":"Testing Task Deskripsi"}`
	request, _ := http.NewRequest("POST", "/create-task", bytes.NewBufferString(testTask))
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "OK response is expected")
}
