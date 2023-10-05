package handlers

import (
	"casual_question/repository"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChannelController struct {
	channelModelRepository repository.ChannelRepositorier
}

func NewChannelController(channelRepo repository.ChannelRepositorier) *ChannelController {
	return &ChannelController{channelModelRepository: channelRepo}
}

func (con ChannelController) MakeChannel(c *gin.Context) {

	// トークンのClaimから得たuserIDを取得
	ownerID := c.GetString("userID")

	id, err := con.channelModelRepository.CreateUnique(ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (con ChannelController) GetAllQuestions(c *gin.Context) {
	channelID := c.Param("id")
	userID := c.GetString("userID")

	ownerID, err := con.channelModelRepository.GetOwnerByChannelID(channelID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ownerID != userID {
		err := errors.New("no permission to see this")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	questions, err := con.channelModelRepository.ReadAllByID(channelID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (con ChannelController) CheckExistence(c *gin.Context) {
	channelID := c.Param("id")

	err := con.channelModelRepository.CheckExistence(channelID)

	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{})
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	return
}
