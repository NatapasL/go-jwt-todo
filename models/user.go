package models

import (
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model

	Username string `sql:"type:VARCHAR(255)"`
	Password string `sql:"type:VARCHAR(255)"`
}

func (m *UserModel) TableName() string {
	return "users"
}
