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

type ExamAPI struct {
	service interfaces.IExamService
}

func NewExamAPI(
	service interfaces.IExamService,
) *ExamAPI {
	return &ExamAPI{
		service: service,
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

	err := p.service.Update(c, examId, &req)
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

func (p *ExamAPI) AddQuestion(c *gin.Context) gohttp.Response {
	examId := utils.StringToUint(c.Param("exam_id"))
	questionID := utils.StringToUint(c.Param("question_id"))
	userID := c.MustGet("userId").(uint)

	err := p.service.AddQuestion(c, examId, questionID, userID)
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

func (p *ExamAPI) DeleteQuestion(c *gin.Context) gohttp.Response {
	examId := utils.StringToUint(c.Param("exam_id"))
	questionID := utils.StringToUint(c.Param("question_id"))
	userID := c.MustGet("userId").(uint)

	err := p.service.DeleteQuestion(c, examId, questionID, userID)
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

func (p *ExamAPI) GetByID(c *gin.Context) gohttp.Response {
	var res serializers.Exam
	userID := c.MustGet("userId").(uint)
	examId := utils.StringToUint(c.Param("id"))
	exam, questions, err := p.service.GetByID(c, examId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, exam)
	utils.Copy(&res.Questions, &questions)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *ExamAPI) Delete(c *gin.Context) gohttp.Response {
	examId := utils.StringToUint(c.Param("id"))
	userID := c.MustGet("userId").(uint)

	err := p.service.Delete(c, examId, userID)
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

func (p *ExamAPI) MoveQuestion(c *gin.Context) gohttp.Response {
	var req serializers.MoveExamReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	err := p.service.MoveQuestion(c, &req)
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
