package groupQuestions

import (
	"errors"
	"net/http"
	"strconv"

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
	c.JSON(http.StatusCreated, serializer.Response())
}

func GroupQuestionUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	groupQuestionModel, err := models.FindGroupQuestion(id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("groupQuestions", errors.New("Invalid slug")))
		return
	}
	groupQuestionModelValidator := NewGroupQuestionModelValidatorFillWith(groupQuestionModel)
	if err := groupQuestionModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	groupQuestionModelValidator.groupQuestionModel.ID = groupQuestionModel.ID
	if err := groupQuestionModel.Update(groupQuestionModelValidator.groupQuestionModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update success"})
}

func GroupQuestionDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteGroupQuestion(id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("groupQuestions", errors.New("Invalid slug")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete success"})
}

func GroupQuestionGetPaging(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")

	page_int, err := strconv.Atoi(page)
	if err != nil || page_int == 0 {
		page_int = 1
	}

	size_int, err := strconv.Atoi(size)
	if err != nil {
		size_int = 20
	}
	offset_int := (page_int - 1) * size_int

	myUserModel := c.MustGet("my_user_model").(models.UserModel)
	groupQuestionModels, modelCount, err := models.FindGroupQuestionPaging(offset_int, size_int, myUserModel.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("groupQuestions", errors.New("Invalid param")))
		return
	}
	serializer := GroupQuestionsSerializer{c, groupQuestionModels}
	c.JSON(http.StatusOK, gin.H{
		"data":   serializer.Response(),
		"paging": common.GetFormatPaging(modelCount, page_int, size_int)})
}

func GroupQuestionGetAll(c *gin.Context) {
	myUserModel := c.MustGet("my_user_model").(models.UserModel)
	groupQuestionModels, err := models.FindGroupQuestionAll(myUserModel.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("groupQuestions", errors.New("Invalid param")))
		return
	}
	serializer := GroupQuestionsSerializer{c, groupQuestionModels}
	c.JSON(http.StatusOK, serializer.Response())
}
