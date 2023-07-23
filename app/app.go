package app

import (
	"github.com/gin-gonic/gin"

	"goshop/app/api"
	"goshop/config"
)

func InitGinEngine(
	userAPI *api.UserAPI,
	groupQuestionAPI *api.GroupQuestionAPI,
) *gin.Engine {
	cfg := config.GetConfig()
	if cfg.Environment == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	api.RegisterAPI(app, userAPI, groupQuestionAPI)
	return app
}
