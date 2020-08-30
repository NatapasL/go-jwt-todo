package persistences

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/NatapasL/go-jwt-todo/forms"
	"github.com/NatapasL/go-jwt-todo/models"
	"github.com/NatapasL/go-jwt-todo/repositories"
)

type postgresUserRepository struct {
	DB *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repositories.UserRepository {
	return &postgresUserRepository{DB: db}
}

func (r *postgresUserRepository) Find(params forms.FindUserParams) (*models.UserModel, error) {
	stubUser := &models.UserModel{Username: "test", Password: "1234"}

	if params.Username != stubUser.Username || params.Password != stubUser.Password {
		return nil, fmt.Errorf("User not found")
	}
	return stubUser, nil
}
