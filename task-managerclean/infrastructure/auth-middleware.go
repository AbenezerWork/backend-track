package infrastructure

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		fmt.Println(authHeader)

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}
		fmt.Println("new", authParts[1])

		claims, err := VerifyJWT(authParts[1], []byte("secret"))

		if err != nil {
			c.JSON(401, gin.H{"error": "Error parsing JWT"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("claims").(*MyCustomClaims)
		if user.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "You are not authorized for this operation"})
			c.Abort()
			return
		}
		c.Next()
	}
}
