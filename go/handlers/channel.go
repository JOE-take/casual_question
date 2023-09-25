package handlers

import (
	"casual_question/models"
	"casual_question/repository"
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
	channel := &models.Channel{}

	err := c.ShouldBindJSON(channel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := con.channelModelRepository.CreateUnique(channel.Owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (con ChannelController) GetAllQuestions(c *gin.Context) {
	channelID := c.Param("id")
	ownerID, err := con.channelModelRepository.GetOwnerByChannelID(channelID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetString("user_id")
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
