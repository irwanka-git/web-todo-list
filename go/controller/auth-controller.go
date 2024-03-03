package controller

import (
	"encoding/json"
	"irwanka/webtodolist/entity"
	"irwanka/webtodolist/helper"
	"irwanka/webtodolist/repository"
	"irwanka/webtodolist/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	authRepository repository.AuthRepository = repository.NewAuthRepository()
	authService    service.AuthService       = service.NewAuthService(authRepository)
)

type AuthController interface {
	SubmitLogin(w http.ResponseWriter, r *http.Request)
}

func NewAuthController() AuthController {
	return &controller{}
}

func (*controller) SubmitLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var credential entity.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: "Username dan Password Wajib Diisi", Status: false})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	errInput := validate.Struct(credential)
	if errInput != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errInput.Error(), Status: false})
		return
	}

	userLogin, errLogin := authService.AuthLogin(&credential)
	if errLogin != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errLogin.Error(), Status: false})
		return
	}

	access_token, errGenToken := helper.CreateJWTTokenLogin(userLogin)

	if errGenToken != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.ResponseMessage{Message: errGenToken.Error()})
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(helper.ResponseData{
		Status:  true,
		Message: "login berhasil",
		Data: map[string]interface{}{
			"access_token":  access_token,
			"email":         userLogin.Email,
			"nama_pengguna": userLogin.NamaPengguna,
			"uuid":          userLogin.UUID,
		},
	})
}
