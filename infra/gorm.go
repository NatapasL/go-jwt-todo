package infra

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/NatapasL/go-jwt-todo/models"
)

var db *gorm.DB

func ConnectToDb() {
	var err error

	dbUrl := os.Getenv("DATABASE_URL")
	db, err = gorm.Open("postgres", dbUrl)
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
