package repository

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type repo struct{}

var (
	db *gorm.DB
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}
	dbOpen, errDBOpen := gorm.Open(mysql.Open(os.Getenv("DSN_MYSQL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(ti)
		},
	})

	if errDBOpen != nil {
		log.Panicln("error Connection Database")
	}
	db = dbOpen
}
