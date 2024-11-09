package middleware

import (
	"awesomeProject2/jwt/authentication"
	"github.com/gin-gonic/gin"
	"log"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(401, gin.H{"error": "token is not active right now"})
		c.Abort()
		return
	}
	claims, msg := authentication.TokenValidation(token)
	log.Println(claims)
	if msg != "" {
		c.JSON(401, gin.H{"error": msg})
		c.Abort()
		return
	}
	c.Next()
}
