package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	jwt.MapClaims
	Email   string `json:"email"`
	User_id string `json:"user_id"`
}

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

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("secret"), nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": "Error parsing JWT"})
			fmt.Println(err, token.Valid)
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			fmt.Println(err, token.Valid)
			c.Abort()
			return
		}
		c.Set("claims", token.Claims.(jwt.MapClaims))
		c.Next()
	}
}
