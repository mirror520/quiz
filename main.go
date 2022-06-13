package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"

	"github.com/mirror520/quiz/model"
	"github.com/mirror520/quiz/persistent/db"
	"github.com/mirror520/quiz/service/quiz"
)

func main() {
	os.Setenv("CONFIGOR_ENV_PREFIX", "QUIZ")

	// 找不到 config.yaml 會自動找 config.exmaple.yaml
	// config.yaml 含機敏資料需隱藏
	configor.Load(&model.Config, "config.yaml")

	router := gin.Default()
	quizV1 := router.Group("/quiz/v1")

	// 注入 comment.Repository
	comments := db.NewCommentRepository()
	svc := quiz.NewService(comments)

	// POST /quiz/v1/comment
	{
		endpoint := quiz.CreateCommentEndpoint(svc)
		handler := quiz.CreateCommentHandler(endpoint)
		quizV1.POST("/comment", handler)
	}

	// GET /quiz/v1/comment/{uuid}
	{
		endpoint := quiz.GetCommentByUUIDEndpoint(svc)
		handler := quiz.GetCommentByUUIDHandler(endpoint)
		quizV1.GET("/comment/:uuid", handler)
	}

	// PUT /quiz/v1/comment/{uuid}
	{
		endpoint := quiz.ModifyCommentByUUIDEndpoint(svc)
		handler := quiz.ModifyCommentByUUIDHandler(endpoint)
		quizV1.PUT("/comment/:uuid", handler)
	}

	// DELETE /quiz/v1/comment/{uuid}
	{
		endpoint := quiz.RemoveCommentByUUIDEndpoint(svc)
		handler := quiz.RemoveCommentByUUIDHandler(endpoint)
		quizV1.DELETE("/comment/:uuid", handler)
	}

	// Resource 記得應該是複數為宜
	// GET    /quiz/v1/comments        Get all comments
	// GET    /quiz/v1/comments/{uuid} Get comments by UUID
	// POST   /quiz/v1/comments        Create a new comment
	// PUT    /quiz/v1/comments/{uuid} Modify comment by UUID
	// DELETE /quiz/v1/comments/{uuid} Remove comment by UUID

	router.Run(":8080")
}
