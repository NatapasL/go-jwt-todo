package infra

import (
	"github.com/gin-gonic/gin"

	"go-jwt-todo/controllers"
	"go-jwt-todo/middlewares"
)

var authenticationController controllers.AuthenticationController
var todoController controllers.TodoController

func mapRoutes(router *gin.Engine) {
	// redis
	redis, err := GetRedisClient()
	if err != nil {
		panic(err)
	}

	// controllers
	authenticationController := controllers.AuthenticationController{
		Redis: redis,
		DB:    "db",
	}
	todoController := controllers.TodoController{DB: "db"}

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
