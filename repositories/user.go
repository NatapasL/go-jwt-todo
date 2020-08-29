package repositories

import (
	"go-jwt-todo/forms"
	"go-jwt-todo/models"
)

type UserRepository interface {
	Find(params forms.FindUserParams) (*models.UserModel, error)
}
