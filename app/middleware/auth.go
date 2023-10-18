package middleware

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"quiztest/pkg/jtoken"
)

func JWTAuth() gin.HandlerFunc {
	return JWT(jtoken.AccessTokenType)
}

func JWTRefresh() gin.HandlerFunc {
	return JWT(jtoken.RefreshTokenType)
}

func JWT(tokenType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		payload, err := jtoken.ValidateToken(token)
		if err != nil || payload == nil || payload["type"] != tokenType {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		c.Set("userId", uint(payload["id"].(float64)))
		c.Set("role", payload["role"])
		// c.Set("userId", uint(1))
		// c.Set("role", "customer")
		c.Next()
	}
}
