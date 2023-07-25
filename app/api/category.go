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

type CategoryAPI struct {
	validator validation.Validation
	service   services.ICategoryService
}

func NewCategoryAPI(
	validator validation.Validation,
	service services.ICategoryService,
) *CategoryAPI {
	return &CategoryAPI{
		validator: validator,
		service:   service,
	}
}

func (p *CategoryAPI) Create(c *gin.Context) gohttp.Response {
	var req serializers.CreateCategoryReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	subjectsReq := []string{}
	for _, item := range req.Subjects {
		subjectsReq = append(subjectsReq, item.Name)
	}

	if !utils.IsUniqueArray(subjectsReq) {
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	req.UserID = utils.StringToUint(c.GetString("userId"))

	category, subjects, err := p.service.Create(c, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.CreateCategoryRes
	utils.Copy(&res, &category)

	var subjectsRes []*serializers.Subject
	utils.Copy(&subjectsRes, &subjects)

	res.Subjects = subjectsRes

	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *CategoryAPI) GetPaging(c *gin.Context) gohttp.Response {
	var req serializers.GetPagingCategoryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = utils.StringToUint(c.GetString("userId"))

	var res serializers.GetPagingCategoryRes

	categories, pagination, err := p.service.GetPaging(c, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res.Categories, &categories)
	res.Pagination = pagination
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *CategoryAPI) GetAll(c *gin.Context) gohttp.Response {
	var res []*serializers.Category
	userID := utils.StringToUint(c.GetString("userId"))
	categories, err := p.service.GetAll(c, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &categories)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *CategoryAPI) Update(c *gin.Context) gohttp.Response {
	categoryId := utils.StringToUint(c.Param("id"))
	var req serializers.UpdateCategoryReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = utils.StringToUint(c.GetString("userId"))

	category, err := p.service.Update(c, categoryId, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Category
	utils.Copy(&res, &category)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *CategoryAPI) GetByID(c *gin.Context) gohttp.Response {
	var res serializers.Category
	userID := utils.StringToUint(c.GetString("userId"))
	categoryId := utils.StringToUint(c.Param("id"))
	category, err := p.service.GetByID(c, categoryId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &category)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *CategoryAPI) Delete(c *gin.Context) gohttp.Response {
	categoryId := utils.StringToUint(c.Param("id"))
	userID := utils.StringToUint(c.GetString("userId"))

	category, err := p.service.Delete(c, categoryId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Category
	utils.Copy(&res, &category)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
