package databases

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Connect() (db *gorm.DB, err error) {
	if err = godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err = gorm.Open("mysql", conn)

	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}
