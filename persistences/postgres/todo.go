package persistences

import (
	"github.com/jinzhu/gorm"

	"github.com/NatapasL/go-jwt-todo/forms"
	"github.com/NatapasL/go-jwt-todo/models"
	"github.com/NatapasL/go-jwt-todo/repositories"
)

type postgresTodoRepository struct {
	DB *gorm.DB
}

func NewPostgresTodoRepository(db *gorm.DB) repositories.TodoRepository {
	return &postgresTodoRepository{DB: db}
}

func (r *postgresTodoRepository) Create(params forms.CreateTodoParams) (*models.TodoModel, error) {
	todo := &models.TodoModel{
		UserID: params.UserID,
		Title:  params.Title,
	}

	r.DB.Create(&todo)

	return todo, nil
}
