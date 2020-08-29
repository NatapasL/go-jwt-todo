package repositories

import (
	"go-jwt-todo/forms"
	"go-jwt-todo/models"
)

type TodoRepository interface {
	Create(params forms.CreateTodoParams) (*models.TodoModel, error)
}
