package handlers

import (
	"casual_question/models"
	"casual_question/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type QuestionController struct {
	questionModelRepository repository.QuestionRepositorier
}

func NewQuestionController(questionRepo repository.QuestionRepositorier) *QuestionController {
	return &QuestionController{questionModelRepository: questionRepo}
}

func (con QuestionController) PostQuestion(c *gin.Context) {
	question := &models.Question{}
	question.ChannelID = c.Param("id")
	question.ID = uuid.New().String()

	// unmarshall
	err := c.ShouldBindJSON(question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// レコード作成
	err = con.questionModelRepository.Create(question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
