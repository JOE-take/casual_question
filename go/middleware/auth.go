package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAccessToken(c *gin.Context) {

	// Bearer で始まっているかチェック
	token := c.GetHeader("Authorization")
	if token[:7] != "Bearer " {
		err := errors.New("header value must start with 'Bearer '")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Next()
}
