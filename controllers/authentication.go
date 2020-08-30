package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"

	"github.com/NatapasL/go-jwt-todo/forms"
	"github.com/NatapasL/go-jwt-todo/persistences/postgres"
	"github.com/NatapasL/go-jwt-todo/services"
)

type AuthenticationController struct {
	Redis *redis.Client
	DB    *gorm.DB
}

func (controller AuthenticationController) Login(c *gin.Context) {
	var params forms.FindUserParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	userRepository := persistences.NewPostgresUserRepository(controller.DB)
	user, err := userRepository.Find(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid username or password")
		return
	}

	authService := services.NewAuthenticationService(controller.Redis)
	tokenDetails, err := authService.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  tokenDetails.AccessToken,
		"refresh_token": tokenDetails.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func (controller AuthenticationController) Logout(c *gin.Context) {
	access, ok := c.Get("access_details")
	if !ok {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	accessDetails, ok := access.(services.AccessDetails)
	if !ok {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	authService := services.NewAuthenticationService(controller.Redis)
	err := authService.DeleteAuth(accessDetails.AccessUuid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func (controller AuthenticationController) Refresh(c *gin.Context) {
	var params forms.RefreshTokenParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := params.RefreshToken

	authService := services.NewAuthenticationService(controller.Redis)
	tokenDetails, err := authService.RefreshAuth(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  tokenDetails.AccessToken,
		"refresh_token": tokenDetails.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}
