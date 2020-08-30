package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/NatapasL/go-jwt-todo/forms"
	"github.com/NatapasL/go-jwt-todo/persistences/postgres"
	"github.com/NatapasL/go-jwt-todo/services"
)

type TodoController interface {
	Create(c *gin.Context)
}

type todoController struct {
	DB *gorm.DB
}

func NewTodoController(db *gorm.DB) TodoController {
	return &todoController{DB: db}
}

func (controller *todoController) Create(c *gin.Context) {
	var params forms.CreateTodoParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	access, ok := c.Get("access_details")
	if !ok {
		c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	accessDetails, ok := access.(services.AccessDetails)
	if !ok {
		c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	params.UserID = accessDetails.UserId

	todoRepository := persistences.NewPostgresTodoRepository(controller.DB)
	todo, err := todoRepository.Create(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, todo)
}
