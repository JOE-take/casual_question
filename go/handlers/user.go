package handlers

import (
	"casual_question/models"
	"casual_question/repository"
	"casual_question/utility"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserController struct {
	userModelRepository repository.UserRepositorier
}

func NewUserController(repo repository.UserRepositorier) *UserController {
	return &UserController{userModelRepository: repo}
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
	user.Password, err = utility.HashingPassword(user.Password)
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
	ok, err := utility.ValidPassword(existingUser.Password, requestUser.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utility.GenerateToken(existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token})
}
