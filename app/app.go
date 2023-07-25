package app

import (
	"github.com/gin-gonic/gin"

	"quiztest/app/api"
	"quiztest/config"
)

func InitGinEngine(
	userAPI *api.UserAPI,
	groupQuestionAPI *api.GroupQuestionAPI,
	categoryAPI *api.CategoryAPI,
	subjectAPI *api.SubjectAPI,
) *gin.Engine {
	cfg := config.GetConfig()
	if cfg.Environment == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	api.RegisterAPI(app, userAPI, groupQuestionAPI, categoryAPI, subjectAPI)
	return app
}