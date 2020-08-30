package infra

import (
	"github.com/gin-gonic/gin"

	"github.com/NatapasL/go-jwt-todo/controllers"
	"github.com/NatapasL/go-jwt-todo/middlewares"
)

var authenticationController controllers.AuthenticationController
var todoController controllers.TodoController

func mapRoutes(router *gin.Engine) {
	// redis
	redis := GetRedisClient()
	db := GetDB()

	ConnectToDb()
	_ = GetDB()

	// controllers
	authenticationController := controllers.NewAuthenticationController(redis, db)
	todoController := controllers.NewTodoController(db)

	// middlewares
	authenticationMiddleware := middlewares.AuthenticationMiddleware(redis)

	api := router.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/login", authenticationController.Login)
		v1.POST("/todo", authenticationMiddleware, todoController.Create)
		v1.POST("/logout", authenticationMiddleware, authenticationController.Logout)
		v1.POST("/token/refresh", authenticationController.Refresh)
	}
}
