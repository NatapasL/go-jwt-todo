package infra

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/NatapasL/go-jwt-todo/models"
)

var db *gorm.DB

func ConnectToDb() {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	var err error

	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	fmt.Println(dbInfo)
	db, err = gorm.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	initModel()
}

func GetDB() *gorm.DB {
	return db
}

func initModel() {
	if !db.HasTable(&models.UserModel{}) {
		db.CreateTable(&models.UserModel{})
	}
	if !db.HasTable(&models.TodoModel{}) {
		db.CreateTable(&models.TodoModel{})
	}
}
