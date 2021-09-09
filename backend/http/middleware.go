package http

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(c *gin.Context) {
	headerToken := c.GetHeader("jwt-token")
	token := strings.Split(headerToken, " ")
	if token[0] != "JWT" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Not Authorized",
		})
		return
	}

	jwtToken, err := jwt.Parse(token[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Not Authorized",
		})
		return
	}

	if !jwtToken.Valid {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Not Authorized",
		})
		return
	}

	c.Next()
}

func verify(c *gin.Context) {
	c.JSON(200, gin.H{})
	return
}
