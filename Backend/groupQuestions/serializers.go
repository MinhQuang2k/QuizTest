package groupQuestions

import (
	"blog.com/models"
	"github.com/gin-gonic/gin"
)

type GroupQuestionSerializer struct {
	C *gin.Context
	models.GroupQuestionModel
}

type GroupQuestionResponse struct {
	ID        uint   `json:"-"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GroupQuestionsSerializer struct {
	C              *gin.Context
	GroupQuestions []models.GroupQuestionModel
}

func (s *GroupQuestionSerializer) Response() GroupQuestionResponse {
	response := GroupQuestionResponse{
		ID:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt: s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
	}
	return response
}

func (s *GroupQuestionsSerializer) Response() []GroupQuestionResponse {
	response := []GroupQuestionResponse{}
	for _, groupQuestion := range s.GroupQuestions {
		serializer := GroupQuestionSerializer{s.C, groupQuestion}
		response = append(response, serializer.Response())
	}
	return response
}
