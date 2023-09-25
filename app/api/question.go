package api

import (
	"quiztest/pkg/errors"
	gohttp "quiztest/pkg/http"
	"quiztest/pkg/logger"

	"github.com/gin-gonic/gin"

	"quiztest/app/interfaces"
	"quiztest/app/serializers"
	"quiztest/pkg/utils"
)

type QuestionAPI struct {
	service interfaces.IQuestionService
}

func NewQuestionAPI(
	service interfaces.IQuestionService,
) *QuestionAPI {
	return &QuestionAPI{
		service: service,
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

	var res serializers.Question

	utils.Copy(&res, &question)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *QuestionAPI) Clones(c *gin.Context) gohttp.Response {
	questionClonesID := utils.StringToUint(c.Param("id"))
	userID := c.MustGet("userId").(uint)

	question, err := p.service.Clones(c, userID, questionClonesID)
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

	err := p.service.Update(c, questionId, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
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

	err := p.service.Delete(c, questionId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
	}
}
