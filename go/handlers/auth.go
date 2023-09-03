package handlers

import (
	"casual_question/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con UserController) Refresh(c *gin.Context) {

	// Bearer で始まっているかチェック
	token := c.GetHeader("Authorization")
	if token[:7] != "Bearer " {
		err := errors.New("header value must start with 'Bearer '")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// リフレッシュトークンのチェック
	refreshToken := token[7:]
	ok, err := util.ValidateRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else if !ok {
		err := errors.New("refresh token isn't valid")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 使われたリフレッシュトークンからユーザ情報を取得
	tokenInfo, err := con.refTokenModelRepository.ReadByToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userInfo, err := con.userModelRepository.ReadByID(tokenInfo.UserID)

	// 新しいアクセストークンを生成
	newAccessToken, err := util.GenerateAccessToken(userInfo)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 新しいリフレッシュトークンを生成
	newRefreshToken, _, err := util.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"AccessToken":  newAccessToken,
		"RefreshToken": newRefreshToken,
	})
}
