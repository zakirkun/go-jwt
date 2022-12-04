package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	helper "github.com/vandenbill/brand-commerce/Helper"
	service "github.com/vandenbill/brand-commerce/Service"
	"net/http"
)

// TODO implement logging here

func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenCookie, _ := c.Cookie("acces_token")
		if tokenCookie == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(tokenCookie, c)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("sub", claims["Sub"])
		} else {
			fmt.Println(err)
		}
	}
}
