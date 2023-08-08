package api

import (
	"quiztest/pkg/errors"
	"quiztest/pkg/logger"

	"github.com/gin-gonic/gin"

	"quiztest/app/interfaces"
	"quiztest/app/serializers"
	"quiztest/pkg/utils"

	gohttp "quiztest/pkg/http"
)

type UserAPI struct {
	service interfaces.IUserService
}

func NewUserAPI(service interfaces.IUserService) *UserAPI {
	return &UserAPI{
		service: service,
	}
}

func (u *UserAPI) Login(c *gin.Context) gohttp.Response {
	var req serializers.LoginReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		return gohttp.Response{
			Error: errors.InvalidParams.Newm(err.Error()),
		}
	}

	user, accessToken, refreshToken, err := u.service.Login(c, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.LoginRes
	utils.Copy(&res.User, &user)
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (u *UserAPI) Register(c *gin.Context) gohttp.Response {
	var req serializers.RegisterReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	user, err := u.service.Register(c, &req)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.RegisterRes
	utils.Copy(&res.User, &user)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (u *UserAPI) GetMe(c *gin.Context) gohttp.Response {
	userID := c.MustGet("userId").(uint)
	user, err := u.service.GetByID(c, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res serializers.User
	utils.Copy(&res, &user)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (u *UserAPI) RefreshToken(c *gin.Context) gohttp.Response {
	userID := c.MustGet("userId").(uint)
	accessToken, err := u.service.RefreshToken(c, userID)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	res := serializers.RefreshTokenRes{
		AccessToken: accessToken,
	}
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}

func (u *UserAPI) ChangePassword(c *gin.Context) gohttp.Response {
	var req serializers.ChangePasswordReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	userID := c.MustGet("userId").(uint)
	err := u.service.ChangePassword(c, userID, &req)
	if err != nil {
		return gohttp.Response{
			Error: err,
		}
	}
	return gohttp.Response{
		Error: errors.Success.New(),
	}
}
