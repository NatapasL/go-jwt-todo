package persistences

import (
	"github.com/jinzhu/gorm"

	"go-jwt-todo/forms"
	"go-jwt-todo/models"
	"go-jwt-todo/repositories"
)

type postgresTodoRepository struct {
	DB *gorm.DB
}

func NewPostgresTodoRepository(db *gorm.DB) repositories.TodoRepository {
	return &postgresTodoRepository{DB: db}
}

func (r *postgresTodoRepository) Create(params forms.CreateTodoParams) (*models.TodoModel, error) {
	todoModel := &models.TodoModel{
		UserID: params.UserID,
		Title:  params.Title,
	}

	return todoModel, nil
}
