package middleware

import (
	"casual_question/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAccessToken(c *gin.Context) {

	// Bearer で始まっているかチェック
	tokenValue := c.GetHeader("Authorization")
	if tokenValue[:7] != "Bearer " {
		err := errors.New("header value must start with 'Bearer '")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// アクセストークンのチェック
	token := tokenValue[7:]
	_, err := util.ParseAccessToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Next()
}
