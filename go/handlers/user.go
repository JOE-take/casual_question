package handlers

import (
	"casual_question/models"
	"casual_question/repository"
	"casual_question/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserController struct {
	userModelRepository     repository.UserRepositorier
	refTokenModelRepository repository.RefTokenRepositorier
}

func NewUserController(userRepo repository.UserRepositorier, refTokenRepo repository.RefTokenRepositorier) *UserController {
	return &UserController{
		userModelRepository:     userRepo,
		refTokenModelRepository: refTokenRepo,
	}
}

func (con UserController) Signup(c *gin.Context) {
	user := &models.User{}

	// unmarshal
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// uuidでUserIDを生成して付与
	user.UserID = uuid.New().String()

	// パスワードのハッシュ化
	user.Password, err = util.HashingPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// レコードの作成
	err = con.userModelRepository.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (con UserController) Login(c *gin.Context) {
	requestUser := &models.User{}

	// unmarshall
	err := c.ShouldBindJSON(requestUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// レコードの読み出し
	existingUser, err := con.userModelRepository.ReadByEmail(requestUser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// メアドチェック
	if existingUser.Email != requestUser.Email {
		err := errors.New("email doesn't match")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードチェック
	ok, err := util.ValidPassword(existingUser.Password, requestUser.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := util.GenerateAccessToken(existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refreshToken, exp, err := util.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// リフレッシュトークンの情報をDBに挿入
	token := &models.RefreshToken{
		Token:  refreshToken,
		UserID: existingUser.UserID,
		Expiry: exp,
	}
	err = con.refTokenModelRepository.Create(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// RefreshTokenをHttpOnlyでCookieに保管
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	// AccessTokenはJSONで返す
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
		"userName":    existingUser.UserName,
		"userId":      existingUser.UserID,
	})
}
