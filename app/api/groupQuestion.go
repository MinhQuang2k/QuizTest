package api

import (
	"net/http"

	"goshop/pkg/logger"
	"goshop/pkg/validation"

	"github.com/gin-gonic/gin"

	"goshop/app/serializers"
	"goshop/app/services"
	"goshop/pkg/response"
	"goshop/pkg/utils"
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

func (p *GroupQuestionAPI) CreateGroupQuestion(c *gin.Context) {
	var req serializers.CreateGroupQuestionReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	req.UserID = c.GetString("userId")

	groupQuestion, err := p.service.Create(c, &req)
	if err != nil {
		logger.Error("Failed to create groupQuestion", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res serializers.GroupQuestion
	utils.Copy(&res, &groupQuestion)
	response.JSON(c, http.StatusOK, res)
}

func (p *GroupQuestionAPI) ListGroupQuestions(c *gin.Context) {
	var req serializers.ListGroupQuestionReq
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	req.UserID = c.GetString("userId")

	var res serializers.ListGroupQuestionRes

	groupQuestions, pagination, err := p.service.ListGroupQuestions(c, &req)
	if err != nil {
		logger.Error("Failed to get list groupQuestions: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.Copy(&res.GroupQuestions, &groupQuestions)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

func (p *GroupQuestionAPI) UpdateGroupQuestion(c *gin.Context) {
	groupQuestionId := c.Param("id")
	var req serializers.UpdateGroupQuestionReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	req.UserID = c.GetString("userId")

	groupQuestion, err := p.service.Update(c, groupQuestionId, &req)
	if err != nil {
		logger.Error("Failed to update groupQuestion", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res serializers.GroupQuestion
	utils.Copy(&res, &groupQuestion)
	response.JSON(c, http.StatusOK, res)
}

func (p *GroupQuestionAPI) GetGroupQuestionByID(c *gin.Context) {
	var res serializers.GroupQuestion
	userID := c.GetString("userId")
	groupQuestionId := c.Param("id")
	groupQuestion, err := p.service.GetGroupQuestionByID(c, groupQuestionId, userID)
	if err != nil {
		logger.Error("Failed to get groupQuestion detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	utils.Copy(&res, &groupQuestion)
	response.JSON(c, http.StatusOK, res)
}
