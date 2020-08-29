package persistences

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"go-jwt-todo/forms"
	"go-jwt-todo/models"
	"go-jwt-todo/repositories"
)

type postgresUserRepository struct {
	DB *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repositories.UserRepository {
	return &postgresUserRepository{DB: db}
}

func (r *postgresUserRepository) Find(params forms.FindUserParams) (*models.UserModel, error) {
	stubUser := &models.UserModel{ID: 1, Username: "test", Password: "1234"}

	if params.Username != stubUser.Username || params.Password != stubUser.Password {
		return nil, fmt.Errorf("User not found")
	}
	return stubUser, nil
}
