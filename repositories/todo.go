package repositories

import (
	"github.com/NatapasL/go-jwt-todo/forms"
	"github.com/NatapasL/go-jwt-todo/models"
)

type TodoRepository interface {
	Create(params forms.CreateTodoParams) (*models.TodoModel, error)
}
