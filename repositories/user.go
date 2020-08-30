package repositories

import (
	"github.com/NatapasL/go-jwt-todo/forms"
	"github.com/NatapasL/go-jwt-todo/models"
)

type UserRepository interface {
	Find(params forms.FindUserParams) (*models.UserModel, error)
}
