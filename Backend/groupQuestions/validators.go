package groupQuestions

import (
	"blog.com/common"
	"blog.com/models"
	"github.com/gin-gonic/gin"
)

type GroupQuestionModelValidator struct {
	GroupQuestion struct {
		Name string `form:"name" json:"name" binding:"max=2048"`
	} `json:"groupQuestion"`
	groupQuestionModel models.GroupQuestionModel `json:"-"`
}

func NewGroupQuestionModelValidator() GroupQuestionModelValidator {
	return GroupQuestionModelValidator{}
}

func NewGroupQuestionModelValidatorFillWith(groupQuestionModel models.GroupQuestionModel) GroupQuestionModelValidator {
	groupQuestionModelValidator := NewGroupQuestionModelValidator()
	groupQuestionModelValidator.GroupQuestion.Name = groupQuestionModel.Name
	return groupQuestionModelValidator
}

func (s *GroupQuestionModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.groupQuestionModel.Name = s.GroupQuestion.Name
	s.groupQuestionModel.User = myUserModel
	return nil
}
