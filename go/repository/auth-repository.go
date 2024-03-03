package repository

import (
	"errors"
	"irwanka/webtodolist/entity"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	AuthLogin(credential *entity.UserCredentials) (*entity.User, error)
	GetUserByUuid(uuid string) (*entity.User, error)
}

func NewAuthRepository() AuthRepository {
	return &repo{}
}

func (*repo) GetUserByUuid(uuid string) (*entity.User, error) {
	var userCek *entity.User
	result := db.Table("users").Where("uuid = ? ", uuid).First(&userCek)
	if result.RowsAffected == 0 {
		return nil, errors.New("email tidak ditemukan")
	}
	return userCek, nil
}

func (*repo) AuthLogin(credential *entity.UserCredentials) (*entity.User, error) {
	var userCek *entity.User

	result := db.Table("users").Where("email = ? ", credential.Email).First(&userCek)

	if result.RowsAffected == 0 {
		return nil, errors.New("email tidak ditemukan")
	}
	errorPassword := bcrypt.CompareHashAndPassword([]byte(userCek.Password), []byte(credential.Password))
	if errorPassword != nil {
		return nil, errors.New("password tidak valid")
	}
	return userCek, nil
}
