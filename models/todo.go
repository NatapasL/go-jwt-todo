package models

import (
	"github.com/jinzhu/gorm"
)

type TodoModel struct {
	gorm.Model

	UserID uint   `sql:"type:VARCHAR(255)"`
	Title  string `sql:"type:VARCHAR(255)"`
}

func (m *TodoModel) TableName() string {
	return "todos"
}
