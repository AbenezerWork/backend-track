package infrastructure

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type MyCustomClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
}

func GenerateJWT(userID, role, email string, expirationTime time.Duration) (string, error) {
	claims := MyCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
		Role:   role,
		Email:  email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("secret")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyJWT(tokenString string, secretKey []byte) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
