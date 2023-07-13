package groupQuestions

import (
	"github.com/gin-gonic/gin"
)

func GroupQuestionRegister(router *gin.RouterGroup) {
	router.GET("/all", GroupQuestionGetAll)
	router.GET("/", GroupQuestionGetPaging)
	router.POST("/", GroupQuestionCreate)
	router.PUT("/:id", GroupQuestionUpdate)
	router.DELETE("/:id", GroupQuestionDelete)
}
