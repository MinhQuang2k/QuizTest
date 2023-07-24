package api

import (
	"github.com/gin-gonic/gin"

	"quiztest/app/middleware"
	"quiztest/pkg/http/wrapper"
)

func RegisterAPI(r *gin.Engine, userAPI *UserAPI, groupQuestionAPI *GroupQuestionAPI, categoryAPI *CategoryAPI, subjectAPI *SubjectAPI) {

	authMiddleware := middleware.JWTAuth()
	refreshAuthMiddleware := middleware.JWTRefresh()
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", wrapper.Wrap(userAPI.Register))
		authRoute.POST("/login", wrapper.Wrap(userAPI.Login))
		authRoute.POST("/refresh", refreshAuthMiddleware, wrapper.Wrap(userAPI.RefreshToken))
		authRoute.GET("/me", authMiddleware, wrapper.Wrap(userAPI.GetMe))
		authRoute.PUT("/change-password", authMiddleware, wrapper.Wrap(userAPI.ChangePassword))
	}

	//--------------------------------API-----------------------------------
	api1 := r.Group("/api")

	// GroupQuestion
	groupQuestionRoute := api1.Group("/group-questions")
	{
		groupQuestionRoute.GET("", authMiddleware, wrapper.Wrap(groupQuestionAPI.GetPaging))
		groupQuestionRoute.GET("/all", authMiddleware, wrapper.Wrap(groupQuestionAPI.GetAll))
		groupQuestionRoute.GET("/:id", authMiddleware, wrapper.Wrap(groupQuestionAPI.GetByID))
		groupQuestionRoute.POST("", authMiddleware, wrapper.Wrap(groupQuestionAPI.Create))
		groupQuestionRoute.PUT("/:id", authMiddleware, wrapper.Wrap(groupQuestionAPI.Update))
		groupQuestionRoute.DELETE("/:id", authMiddleware, wrapper.Wrap(groupQuestionAPI.Delete))
	}

	// Category
	categoryRoute := api1.Group("/categories")
	{
		categoryRoute.GET("", authMiddleware, wrapper.Wrap(categoryAPI.GetPaging))
		categoryRoute.GET("/all", authMiddleware, wrapper.Wrap(categoryAPI.GetAll))
		categoryRoute.GET("/:id", authMiddleware, wrapper.Wrap(categoryAPI.GetByID))
		categoryRoute.POST("", authMiddleware, wrapper.Wrap(categoryAPI.Create))
		categoryRoute.PUT("/:id", authMiddleware, wrapper.Wrap(categoryAPI.Update))
		categoryRoute.DELETE("/:id", authMiddleware, wrapper.Wrap(categoryAPI.Delete))
	}

	// Subject
	subjectRoute := api1.Group("/subjects")
	{
		subjectRoute.GET("/all", authMiddleware, wrapper.Wrap(subjectAPI.GetAll))
		subjectRoute.POST("", authMiddleware, wrapper.Wrap(subjectAPI.Create))
		subjectRoute.PUT("/:id", authMiddleware, wrapper.Wrap(subjectAPI.Update))
		subjectRoute.DELETE("/:id", authMiddleware, wrapper.Wrap(subjectAPI.Delete))
	}

}
