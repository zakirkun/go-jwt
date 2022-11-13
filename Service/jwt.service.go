package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string, ctx *gin.Context) (*jwt.Token, error)
}

type jwtServiceImpl struct {
	secretKey string
}

func NewJwtServiceImpl(secretKey string) JwtService {
	return &jwtServiceImpl{secretKey: secretKey}
}

func (j *jwtServiceImpl) GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Sub":       userID,
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": time.Now().AddDate(0, 0, 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(j.secretKey))

	return tokenString, err
}

func (j *jwtServiceImpl) ValidateToken(tokenString string, ctx *gin.Context) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})

	return token, err
}
