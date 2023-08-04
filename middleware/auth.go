package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func ValidateJWT(c *gin.Context) (*jwt.Token, error) {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can not load .env file", err)
	}

	header := c.Request.Header.Get("Authorization")

	tokenString := strings.TrimPrefix(header, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return token, err
}

func UserAuth(c *gin.Context) {

	token, err := ValidateJWT(c)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("username", claims["username"])
	} else {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "forbidden", "message": err.Error()})
		return
	}
	c.Next()
}

func AdminAuth(c *gin.Context) {

	token, err := ValidateJWT(c)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["role"] == "admin" {
			c.Set("username", claims["username"])
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "forbidden", "message": "You have no permission"})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "forbidden", "message": err.Error()})
		return
	}
	c.Next()
}
