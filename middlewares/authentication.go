package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"

	"github.com/NatapasL/go-jwt-todo/services"
)

func AuthenticationMiddleware(r *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c.Request)
		verifyTokenService := services.NewAuthenticationService(r)
		accessDetails, err := verifyTokenService.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		c.Set("access_details", *accessDetails)
		c.Next()
	}
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
