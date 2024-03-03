package main

import (
	"irwanka/webtodolist/controller"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"

	customMiddleware "irwanka/webtodolist/middleware"
)

var (
	tokenAuth      *jwtauth.JWTAuth
	homeController controller.HomeController       = controller.NewHomeController()
	authController controller.AuthController       = controller.NewAuthController()
	taskController controller.TaskController       = controller.NewTaskController()
	userMiddleware customMiddleware.UserMiddleware = customMiddleware.NewUserMiddleware()
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	signKey := os.Getenv("JWT_SIGN_KEY")
	tokenAuth = jwtauth.New("HS512", []byte(signKey), nil)
}

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	port := os.Getenv("PORT")
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(120 * time.Second))

	//http://localhost:8000/api
	r.Route("/api", func(r chi.Router) {
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

	http.ListenAndServe(port, r)

}
