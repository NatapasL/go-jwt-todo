package infra

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"go-jwt-todo/models"
)

var db *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "test_test"
)

func ConnectToDb() {
	var err error

	dbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
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
