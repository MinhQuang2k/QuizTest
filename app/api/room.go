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

type RoomAPI struct {
	validator validation.Validation
	service   services.IRoomService
}

func NewRoomAPI(
	validator validation.Validation,
	service services.IRoomService,
) *RoomAPI {
	return &RoomAPI{
		validator: validator,
		service:   service,
	}
}

func (p *RoomAPI) Create(c *gin.Context) gohttp.Response {
	var req serializers.CreateRoomReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	room, err := p.service.Create(c, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Room
	utils.Copy(&res, &room)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *RoomAPI) GetPaging(c *gin.Context) gohttp.Response {
	var req serializers.GetPagingRoomReq
	if err := c.ShouldBindQuery(&req); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	var res serializers.GetPagingRoomRes

	rooms, pagination, err := p.service.GetPaging(c, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res.Rooms, &rooms)
	res.Pagination = pagination
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *RoomAPI) Update(c *gin.Context) gohttp.Response {
	roomId := utils.StringToUint(c.Param("id"))
	var req serializers.UpdateRoomReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	req.UserID = c.MustGet("userId").(uint)

	room, err := p.service.Update(c, roomId, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Room
	utils.Copy(&res, &room)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *RoomAPI) GetByID(c *gin.Context) gohttp.Response {
	var res serializers.Room
	userID := c.MustGet("userId").(uint)
	roomId := utils.StringToUint(c.Param("id"))
	room, err := p.service.GetByID(c, roomId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	utils.Copy(&res, &room)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (p *RoomAPI) GetCodeRoom(c *gin.Context) gohttp.Response {
	var req serializers.GetCodeRoomReq
	if err := c.ShouldBindQuery(&req); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	var listCode []string

	for i := 0; i < req.Limit; i++ {
		code, err := utils.GenerateRandomString(3)
		if err != nil {
			return gohttp.Response{
				Error: errors.Error.New(),
			}
		}
		listCode = append(listCode, code)
	}

	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  listCode,
	}
}

func (p *RoomAPI) Delete(c *gin.Context) gohttp.Response {
	roomId := utils.StringToUint(c.Param("id"))
	userID := c.MustGet("userId").(uint)

	room, err := p.service.Delete(c, roomId, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.Room
	utils.Copy(&res, &room)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
