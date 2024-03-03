package helper

import (
	"irwanka/webtodolist/entity"
	"log"
	"os"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
)

func CreateJWTTokenLogin(user *entity.User) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	signKey := os.Getenv("JWT_SIGN_KEY")

	tokenAuth := jwtauth.New("HS512", []byte(signKey), nil)

	sub := user.UUID
	claims := map[string]interface{}{
		"id":            user.ID,
		"sub":           sub,
		"email":         user.Email,
		"nama_pengguna": user.NamaPengguna,
	}
	jwtauth.SetExpiryIn(claims, time.Hour*12)
	jwtauth.SetIssuedAt(claims, time.Now())
	_, tokenString, _ := tokenAuth.Encode(claims)
	return tokenString, nil
}
