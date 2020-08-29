package persistences

import (
	"go-jwt-todo/forms"
	"go-jwt-todo/models"
	"go-jwt-todo/repositories"
)

type postgresTodoRepository struct {
	DB string
}

func NewPostgresTodoRepository(db string) repositories.TodoRepository {
	return &postgresTodoRepository{DB: db}
}

func (r *postgresTodoRepository) Create(params forms.CreateTodoParams) (*models.TodoModel, error) {
	todoModel := &models.TodoModel{
		UserID: params.UserID,
		Title:  params.Title,
	}

	return todoModel, nil
}
