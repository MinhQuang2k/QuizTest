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

type GroupQuestionAPI struct {
	validator validation.Validation
	service   services.IGroupQuestionService
}

func NewGroupQuestionAPI(
	validator validation.Validation,
	service services.IGroupQuestionService,
) *GroupQuestionAPI {
	return &GroupQuestionAPI{
		validator: validator,
		service:   service,
	}
}

func (p *GroupQuestionAPI) Create(c *gin.Context) gohttp.Response {
	var req serializers.CreateGroupQuestionReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = utils.StringToUint(c.GetString("userId"))

	groupQuestion, err := p.service.Create(c, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.GroupQuestion
	utils.Copy(&res, &groupQuestion)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *GroupQuestionAPI) GetPaging(c *gin.Context) gohttp.Response {
	var req serializers.GetPagingGroupQuestionReq
	if err := c.ShouldBindQuery(&req); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = utils.StringToUint(c.GetString("userId"))

	var res serializers.GetPagingGroupQuestionRes

	groupQuestions, pagination, err := p.service.GetPaging(c, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res.GroupQuestions, &groupQuestions)
	res.Pagination = pagination
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *GroupQuestionAPI) GetAll(c *gin.Context) gohttp.Response {
	var res []*serializers.GroupQuestion
	userID := utils.StringToUint(c.GetString("userId"))
	groupQuestions, err := p.service.GetAll(c, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &groupQuestions)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *GroupQuestionAPI) Update(c *gin.Context) gohttp.Response {
	groupQuestionId := utils.StringToUint(c.Param("id"))
	var req serializers.UpdateGroupQuestionReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = utils.StringToUint(c.GetString("userId"))

	groupQuestion, err := p.service.Update(c, groupQuestionId, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.GroupQuestion
	utils.Copy(&res, &groupQuestion)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *GroupQuestionAPI) GetByID(c *gin.Context) gohttp.Response {
	var res serializers.GroupQuestion
	userID := utils.StringToUint(c.GetString("userId"))
	groupQuestionId := utils.StringToUint(c.Param("id"))
	groupQuestion, err := p.service.GetByID(c, groupQuestionId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &groupQuestion)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *GroupQuestionAPI) Delete(c *gin.Context) gohttp.Response {
	groupQuestionId := utils.StringToUint(c.Param("id"))
	userID := utils.StringToUint(c.GetString("userId"))

	groupQuestion, err := p.service.Delete(c, groupQuestionId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.GroupQuestion
	utils.Copy(&res, &groupQuestion)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
