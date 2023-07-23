package api

import (
	"github.com/gin-gonic/gin"

	"goshop/app/middleware"
)

func RegisterAPI(r *gin.Engine, userAPI *UserAPI, groupQuestionAPI *GroupQuestionAPI) {

	authMiddleware := middleware.JWTAuth()
	refreshAuthMiddleware := middleware.JWTRefresh()
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", userAPI.Register)
		authRoute.POST("/login", userAPI.Login)
		authRoute.POST("/refresh", refreshAuthMiddleware, userAPI.RefreshToken)
		authRoute.GET("/me", authMiddleware, userAPI.GetMe)
		authRoute.PUT("/change-password", authMiddleware, userAPI.ChangePassword)
	}

	//--------------------------------API-----------------------------------
	api1 := r.Group("/api")

	// GroupQuestion
	groupQuestionRoute := api1.Group("/group-questions")
	{
		groupQuestionRoute.GET("", authMiddleware, groupQuestionAPI.ListGroupQuestions)
		groupQuestionRoute.POST("", authMiddleware, groupQuestionAPI.CreateGroupQuestion)
		groupQuestionRoute.PUT("/:id", authMiddleware, groupQuestionAPI.UpdateGroupQuestion)
		groupQuestionRoute.GET("/:id", authMiddleware, groupQuestionAPI.GetGroupQuestionByID)
	}

}
