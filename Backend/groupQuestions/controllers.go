package groupQuestions

import (
	"errors"
	"net/http"

	"blog.com/common"
	"blog.com/models"
	"github.com/gin-gonic/gin"
)

func GroupQuestionCreate(c *gin.Context) {
	groupQuestionModelValidator := NewGroupQuestionModelValidator()
	if err := groupQuestionModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := models.SaveOne(&groupQuestionModelValidator.groupQuestionModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := GroupQuestionSerializer{c, groupQuestionModelValidator.groupQuestionModel}
	c.JSON(http.StatusCreated, gin.H{"groupQuestion": serializer.Response()})
}

// func GroupQuestionUpdate(c *gin.Context) {
// 	slug := c.Param("slug")
// 	groupQuestionModel, err := models.FindGroupQuestion(&models.GroupQuestionModel{Slug: slug})
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, common.NewError("groupQuestions", errors.New("Invalid slug")))
// 		return
// 	}
// 	groupQuestionModelValidator := NewGroupQuestionModelValidatorFillWith(groupQuestionModel)
// 	if err := groupQuestionModelValidator.Bind(c); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
// 		return
// 	}

// 	groupQuestionModelValidator.groupQuestionModel.ID = groupQuestionModel.ID
// 	if err := groupQuestionModel.Update(groupQuestionModelValidator.groupQuestionModel); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
// 		return
// 	}
// 	serializer := GroupQuestionSerializer{c, groupQuestionModel}
// 	c.JSON(http.StatusOK, gin.H{"groupQuestion": serializer.Response()})
// }

// func GroupQuestionDelete(c *gin.Context) {
// 	slug := c.Param("slug")
// 	err := models.DeleteGroupQuestion(&models.GroupQuestionModel{Slug: slug})
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, common.NewError("groupQuestions", errors.New("Invalid slug")))
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"groupQuestion": "Delete success"})
// }

func GroupQuestionGetPaging(c *gin.Context) {
	all := c.Query("all")
	limit := c.Query("limit")
	offset := c.Query("offset")
	myUserModel := c.MustGet("myUserModel").(models.UserModel)
	groupQuestionModels, modelCount, err := models.FindGroupQuestionPaging(all, limit, offset, myUserModel.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("groupQuestions", errors.New("Invalid param")))
		return
	}
	serializer := GroupQuestionsSerializer{c, groupQuestionModels}
	c.JSON(http.StatusOK, gin.H{"groupQuestions": serializer.Response(), "groupQuestionsCount": modelCount})
}
