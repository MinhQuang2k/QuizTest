package api

import (
	"quiztest/pkg/errors"
	gohttp "quiztest/pkg/http"
	"quiztest/pkg/logger"
	"quiztest/pkg/validation"

	"github.com/gin-gonic/gin"

	"quiztest/app/serializers"
	"quiztest/app/services"
	"quiztest/pkg/utils"
)

type QuestionAPI struct {
	validator validation.Validation
	service   services.IQuestionService
}

func NewQuestionAPI(
	validator validation.Validation,
	service services.IQuestionService,
) *QuestionAPI {
	return &QuestionAPI{
		validator: validator,
		service:   service,
	}
}

func (p *QuestionAPI) Create(c *gin.Context) gohttp.Response {
	var req serializers.CreateQuestionReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	question, err := p.service.Create(c, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Question
	utils.Copy(&res, &question)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *QuestionAPI) Clones(c *gin.Context) gohttp.Response {
	questionClonesID := utils.StringToUint(c.Param("id"))
	userID := c.MustGet("userId").(uint)

	err := p.service.Clones(c, userID, questionClonesID)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
	}
}

func (p *QuestionAPI) GetPaging(c *gin.Context) gohttp.Response {
	var req serializers.GetPagingQuestionReq
	if err := c.ShouldBindQuery(&req); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	questions, pagination, err := p.service.GetPaging(c, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.GetPagingQuestionRes

	utils.Copy(&res.Questions, &questions)
	res.Pagination = pagination
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *QuestionAPI) Update(c *gin.Context) gohttp.Response {
	questionId := utils.StringToUint(c.Param("id"))
	var req serializers.UpdateQuestionReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	question, err := p.service.Update(c, questionId, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Question
	utils.Copy(&res, &question)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *QuestionAPI) GetByID(c *gin.Context) gohttp.Response {
	var res serializers.Question
	userID := c.MustGet("userId").(uint)
	questionId := utils.StringToUint(c.Param("id"))
	question, err := p.service.GetByID(c, questionId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &question)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *QuestionAPI) Delete(c *gin.Context) gohttp.Response {
	questionId := utils.StringToUint(c.Param("id"))
	userID := c.MustGet("userId").(uint)

	question, err := p.service.Delete(c, questionId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Question
	utils.Copy(&res, &question)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
