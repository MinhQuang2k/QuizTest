package api

import (
	"quiztest/config"
	"quiztest/pkg/errors"
	gohttp "quiztest/pkg/http"
	"quiztest/pkg/logger"

	"github.com/gin-gonic/gin"

	"quiztest/app/interfaces"
	"quiztest/app/serializers"
	"quiztest/pkg/utils"
)

type CategoryAPI struct {
	service interfaces.ICategoryService
	rdbs    interfaces.ICache
}

func NewCategoryAPI(
	service interfaces.ICategoryService,
	rdbs interfaces.ICache,
) *CategoryAPI {
	return &CategoryAPI{
		service: service,
		rdbs:    rdbs,
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

	req.UserID = c.MustGet("userId").(uint)

	err := p.service.Create(c, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}
	_ = p.cache.RemovePattern("*categories*")
	return gohttp.Response{
		Error: errors.Success.New(),
	}
}

func (p *CategoryAPI) GetPaging(c *gin.Context) gohttp.Response {
	var req serializers.GetPagingCategoryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

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
	err := p.rdbs.GetInstance().Get(c.Request.URL.Path, &res)
	if err == nil {
		return gohttp.Response{
			Error: errors.Success.New(),
			Data:  res,
		}
	}
	userID := c.MustGet("userId").(uint)
	categories, err := p.service.GetAll(c, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &categories)
	_ = p.rdbs.GetInstance().SetWithExpiration(c.Request.URL.Path, res, config.ProductCachingTime)
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

	req.UserID = c.MustGet("userId").(uint)

	err := p.service.Update(c, categoryId, &req)
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

func (p *CategoryAPI) GetByID(c *gin.Context) gohttp.Response {
	var res serializers.Category
	userID := c.MustGet("userId").(uint)
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
	userID := c.MustGet("userId").(uint)

	err := p.service.Delete(c, categoryId, userID)
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
