package router

import (
	"casual_question/config"
	"casual_question/handlers"
	"casual_question/middleware"
	"casual_question/repository"
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(config.CorsConfig()))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "page not found")
	})

	userRepository := repository.NewUserRepository(db)
	refTokenRepository := repository.NewRefRepository(db)
	questionRepository := repository.NewQuestionRepository(db)
	channelRepository := repository.NewChannelRepository(db)
	userController := handlers.NewUserController(userRepository, refTokenRepository)
	questionController := handlers.NewQuestionController(questionRepository)
	channelController := handlers.NewChannelController(channelRepository)

	r.POST("/signup", userController.Signup)
	r.POST("/login", userController.Login)
	r.GET("/refresh", userController.Refresh)

	restricted := r.Group("")
	{
		restricted.Use(middleware.CheckAccessToken)
		restricted.POST("/channel/new", channelController.MakeChannel)
		restricted.GET("/channel/:id", channelController.GetAllQuestions)
	}

	r.POST("/channel/:id/post", questionController.PostQuestion)

	return r
}
