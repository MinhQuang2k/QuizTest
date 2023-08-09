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

type GroupQuestionAPI struct {
	service interfaces.IGroupQuestionService
}

func NewGroupQuestionAPI(
	service interfaces.IGroupQuestionService,
) *GroupQuestionAPI {
	return &GroupQuestionAPI{
		service: service,
	}
}

func (p *GroupQuestionAPI) Create(c *gin.Context) gohttp.Response {
	var req serializers.CreateGroupQuestionReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	err := p.service.Create(c, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
	}
}

func (p *GroupQuestionAPI) GetPaging(c *gin.Context) gohttp.Response {
	var req serializers.GetPagingGroupQuestionReq
	if err := c.ShouldBindQuery(&req); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

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
	userID := c.MustGet("userId").(uint)
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

	req.UserID = c.MustGet("userId").(uint)

	err := p.service.Update(c, groupQuestionId, &req)
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

func (p *GroupQuestionAPI) GetByID(c *gin.Context) gohttp.Response {
	var res serializers.GroupQuestion
	userID := c.MustGet("userId").(uint)
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
	userID := c.MustGet("userId").(uint)

	err := p.service.Delete(c, groupQuestionId, userID)
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
