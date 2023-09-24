package middleware

import (
	"casual_question/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAccessToken(c *gin.Context) {

	// トークンのバリデーション
	tokenValue := c.GetHeader("Authorization")
	if len(tokenValue) < 7 {
		err := errors.New("wrong token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else if tokenValue[:7] != "Bearer " {
		err := errors.New("header value must start with 'Bearer '")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// アクセストークンのチェック
	token := tokenValue[7:]
	claims, err := util.ParseAccessToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("user_id", claims.UserID)
	c.Next()
}
