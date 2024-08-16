package Infrastructure

import (
	"net/http"
	"strings"

	usecases "github.com/Abzaek/clean-arch/Usecases"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Auth usecases.AuthService
}

func (am *AuthMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "header is required"})
			c.Abort()
			return
		}
		authParts := strings.Split(authHeader, " ")

		if strings.ToLower(authParts[0]) != "bearer" || len(authParts) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid header"})
			c.Abort()
			return
		}

		claim, err := am.Auth.ValidateToken(authParts[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("Claims", claim)
		c.Next()
	}
}
