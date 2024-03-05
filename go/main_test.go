package main_test

import (
	"bytes"
	"irwanka/webtodolist/controller"
	customMiddleware "irwanka/webtodolist/middleware"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	tokenAuth      *jwtauth.JWTAuth
	homeController controller.HomeController       = controller.NewHomeController()
	authController controller.AuthController       = controller.NewAuthController()
	taskController controller.TaskController       = controller.NewTaskController()
	userMiddleware customMiddleware.UserMiddleware = customMiddleware.NewUserMiddleware()

	tokenJWTUser1 string = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQGVtYWlsLmNvbSIsImlhdCI6MTcwOTYxOTk0MCwiaWQiOjEsIm5hbWFfcGVuZ2d1bmEiOiJJcndhbiBLdXJuaWF3YW4iLCJzdWIiOiJhMmFhMmNkMC02ZGY1LTRhYmEtYmE0YS0xZTMwZmE5ZGM2NzUifQ.JP49LvJ_wfofm9D6NTj2xIuGeqzUvbJwhBdxWaoU7hiIkrIqOQ_ZIpW1PuoM6p1pAxwKlB3u8AdMI-WpEaxxEg"
	tokenJWTUser2 string = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIyQGVtYWlsLmNvbSIsImlhdCI6MTcwOTYxOTk2MiwiaWQiOjIsIm5hbWFfcGVuZ2d1bmEiOiJBcmlmIE11aGFtbWFkIiwic3ViIjoiMWRmODI3YWItNTQ3YS00Y2ZhLWI2YjItZmVlYWM0MDQwZWM3In0.m8Nmv5EqA5odryj6T45E4o-95Q5L7-lJXsJq56TRhc4nVWUXigF8XSYPKhA1ZF4n__vxpvXISOJSZ2Ck3udbYA"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	signKey := os.Getenv("JWT_SIGN_KEY")
	tokenAuth = jwtauth.New("HS512", []byte(signKey), nil)
}

func Router() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", homeController.Welcome)
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

func TestWelcome(t *testing.T) {
	t.Logf("Test Welcome Response")
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestLogin(t *testing.T) {
	t.Logf("Test Login")
	testUser := `{"email":"user1@email.com" , "password":"123456"}`
	request, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(testUser))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "OK response is expected")
}

func TestCreateTask(t *testing.T) {
	t.Logf("Test Create Task")

	var bearer = "Bearer " + tokenJWTUser1
	testTask := `{"title":"Testing Task" , "description":"Testing Task Deskripsi"}`
	request, _ := http.NewRequest("POST", "/create-task", bytes.NewBufferString(testTask))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "OK response is expected")
}

func TestUpdateTask(t *testing.T) {
	t.Logf("Test Update Task")

	var bearer = "Bearer " + tokenJWTUser1
	var id_task = "3"
	testTask := `{"title":"Testing Task Update" , "description":"Testing Task Deskripsi Update"}`
	request, _ := http.NewRequest("PATCH", "/update-task/"+id_task, bytes.NewBufferString(testTask))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "OK response is expected")
}

func TestGetListTask(t *testing.T) {
	t.Logf("Test List Task")

	var bearer = "Bearer " + tokenJWTUser1
	request, _ := http.NewRequest("GET", "/list-task", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestGetDetilTask(t *testing.T) {
	t.Logf("Test Get Detil Task by ID")

	var bearer = "Bearer " + tokenJWTUser1
	var id_task = "3"
	request, _ := http.NewRequest("GET", "/get-detil-task/"+id_task, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestGetDetilTaskNotAuthorization(t *testing.T) {
	t.Logf("Test Get Detil Task by ID")

	var bearer = "Bearer " + tokenJWTUser2
	var id_task = "3"
	request, _ := http.NewRequest("GET", "/get-detil-task/"+id_task, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 403, response.Code, "OK response is expected")
}
