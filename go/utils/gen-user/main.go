package main

import (
	"fmt"
	"irwanka/webtodolist/entity"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db_akses *gorm.DB
)

func init() {
	dsn := "root:112w4nka@tcp(host.docker.internal:8036)/db_todolist?charset=utf8mb4&parseTime=True&loc=Local"
	loggerGorm := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Minute,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,         // Don't include params in the SQL log
			Colorful:                  false,        // Disable color
		},
	)

	dbOpen2, errDBOpen := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerGorm,
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(ti)
		},
	})

	if errDBOpen != nil {
		log.Panicln("error Connection Database")
	}
	db_akses = dbOpen2
}

func main() {
	fmt.Println("Buat User")
	var user = entity.User{}
	user.Email = "akbarfajar.email@gmail.com"
	user.NamaPengguna = "Akbar Fajar"
	user.UUID = uuid.New().String()
	bytesPassword, errCryptPassword := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	if errCryptPassword != nil {
		log.Panicln("Error loading generate password")
	}
	user.Password = string(bytesPassword)
	tx := db_akses.Begin()
	tx.Table("users").Create(&user)
	tx.Commit()
}
