package handlers

import (
	"casual_question/models"
	"casual_question/repository"
	"casual_question/utility"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// uuidでUserIDを生成して付与
	user.UserID = uuid.New().String()

	// パスワードのハッシュ化
	user.Password, err = utility.HashingPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// レコードの作成
	err = con.userModelRepository.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (con UserController) Login(c *gin.Context) {
	requestUser := &models.User{}
	err := c.ShouldBindJSON(requestUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	existingUser, err := con.userModelRepository.ReadByEmail(requestUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if existingUser.Email != requestUser.Email {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if ok, err := utility.ValidPassword(existingUser.Password, requestUser.Password); ok && existingUser.Email == requestUser.Email {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ok, err := utility.ValidPassword(existingUser.Password, requestUser.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
