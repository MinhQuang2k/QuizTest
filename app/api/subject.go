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

type SubjectAPI struct {
	validator validation.Validation
	service   services.ISubjectService
}

func NewSubjectAPI(
	validator validation.Validation,
	service services.ISubjectService,
) *SubjectAPI {
	return &SubjectAPI{
		validator: validator,
		service:   service,
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

func (p *SubjectAPI) Delete(c *gin.Context) gohttp.Response {
	subjectId := utils.StringToUint(c.Param("id"))
	userID := c.MustGet("userId").(uint)

	subject, err := p.service.Delete(c, subjectId, userID)
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
