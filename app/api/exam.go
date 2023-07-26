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

type ExamAPI struct {
	validator validation.Validation
	service   services.IExamService
}

func NewExamAPI(
	validator validation.Validation,
	service services.IExamService,
) *ExamAPI {
	return &ExamAPI{
		validator: validator,
		service:   service,
	}
}

func (p *ExamAPI) Create(c *gin.Context) gohttp.Response {
	var req serializers.CreateExamReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	exam, err := p.service.Create(c, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Exam
	utils.Copy(&res, &exam)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *ExamAPI) GetPaging(c *gin.Context) gohttp.Response {
	var req serializers.GetPagingExamReq
	if err := c.ShouldBindQuery(&req); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	var res serializers.GetPagingExamRes

	exams, pagination, err := p.service.GetPaging(c, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res.Exams, &exams)
	res.Pagination = pagination
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *ExamAPI) GetAll(c *gin.Context) gohttp.Response {
	var res []*serializers.Exam
	userID := c.MustGet("userId").(uint)
	exams, err := p.service.GetAll(c, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &exams)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *ExamAPI) Update(c *gin.Context) gohttp.Response {
	examId := utils.StringToUint(c.Param("id"))
	var req serializers.UpdateExamReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	exam, err := p.service.Update(c, examId, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Exam
	utils.Copy(&res, &exam)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *ExamAPI) GetByID(c *gin.Context) gohttp.Response {
	var res serializers.Exam
	userID := c.MustGet("userId").(uint)
	examId := utils.StringToUint(c.Param("id"))
	exam, err := p.service.GetByID(c, examId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &exam)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *ExamAPI) Delete(c *gin.Context) gohttp.Response {
	examId := utils.StringToUint(c.Param("id"))
	userID := c.MustGet("userId").(uint)

	exam, err := p.service.Delete(c, examId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Exam
	utils.Copy(&res, &exam)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
