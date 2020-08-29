package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go-jwt-todo/forms"
	"go-jwt-todo/persistences/postgres"
	"go-jwt-todo/services"
)

type TodoController struct {
	DB *gorm.DB
}

func (controller *TodoController) Create(c *gin.Context) {
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
