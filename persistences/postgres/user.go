package persistences

import (
	"fmt"

	"go-jwt-todo/forms"
	"go-jwt-todo/models"
	"go-jwt-todo/repositories"
)

type postgresUserRepository struct {
	DB string
}

func NewPostgresUserRepository(db string) repositories.UserRepository {
	return &postgresUserRepository{DB: db}
}

func (r *postgresUserRepository) Find(params forms.FindUserParams) (*models.UserModel, error) {
	stubUser := &models.UserModel{ID: 1, Username: "test", Password: "1234"}

	if params.Username != stubUser.Username || params.Password != stubUser.Password {
		return nil, fmt.Errorf("User not found")
	}
	return stubUser, nil
}
