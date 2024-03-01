package service

import (
	"irwanka/webtodolist/entity"
	"irwanka/webtodolist/repository"
)

type AuthService interface {
	AuthLogin(credential *entity.UserCredentials) (*entity.User, error)
}

var (
	authRepo repository.AuthRepository
)

func NewAuthService(repo repository.AuthRepository) AuthService {
	authRepo = repo
	return &service{}
}

func (*service) AuthLogin(credential *entity.UserCredentials) (*entity.User, error) {
	return authRepo.AuthLogin(credential)
}
