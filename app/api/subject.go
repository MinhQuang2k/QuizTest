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

type SubjectAPI struct {
	service interfaces.ISubjectService
}

func NewSubjectAPI(
	service interfaces.ISubjectService,
) *SubjectAPI {
	return &SubjectAPI{
		service: service,
	}
}

func (p *SubjectAPI) Create(c *gin.Context) gohttp.Response {
	var req serializers.CreateSubjectReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	subject, err := p.service.Create(c, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Subject
	utils.Copy(&res, &subject)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *SubjectAPI) Update(c *gin.Context) gohttp.Response {
	subjectId := utils.StringToUint(c.Param("id"))
	var req serializers.UpdateSubjectReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	subject, err := p.service.Update(c, subjectId, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Subject
	utils.Copy(&res, &subject)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *SubjectAPI) Move(c *gin.Context) gohttp.Response {
	subjectId := utils.StringToUint(c.Param("id"))
	var req serializers.MoveSubjectReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	subject, err := p.service.Move(c, subjectId, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Subject
	utils.Copy(&res, &subject)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *SubjectAPI) Delete(c *gin.Context) gohttp.Response {
	subjectId := utils.StringToUint(c.Param("id"))
	categoryID := utils.StringToUint(c.Param("category_id"))
	userID := c.MustGet("userId").(uint)

	subject, err := p.service.Delete(c, subjectId, categoryID, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Subject
	utils.Copy(&res, &subject)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
