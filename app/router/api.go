package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"quiztest/app/api"
	"quiztest/app/middleware"
	"quiztest/pkg/http/wrapper"
	"quiztest/pkg/logger"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		userAPI *api.UserAPI,
		groupQuestionAPI *api.GroupQuestionAPI,
		categoryAPI *api.CategoryAPI,
		subjectAPI *api.SubjectAPI,
		questionAPI *api.QuestionAPI,
		examAPI *api.ExamAPI,
		roomAPI *api.RoomAPI,
	) error {
		authMiddleware := middleware.JWTAuth()
		refreshAuthMiddleware := middleware.JWTRefresh()
		authRoute := r.Group("/auth")
		{
			authRoute.POST("/register", wrapper.Wrap(userAPI.Register))
			authRoute.POST("/login", wrapper.Wrap(userAPI.Login))
			authRoute.POST("/refresh", refreshAuthMiddleware, wrapper.Wrap(userAPI.RefreshToken))
			authRoute.GET("/me", authMiddleware, wrapper.Wrap(userAPI.GetMe))
			authRoute.PUT("/change-password", authMiddleware, wrapper.Wrap(userAPI.ChangePassword))
			authRoute.POST("/send-mail", wrapper.Wrap(userAPI.SendMail))
		}

		//--------------------------------API-----------------------------------
		api1 := r.Group("/api")

		// GroupQuestion
		groupQuestionRoute := api1.Group("/group-questions", authMiddleware)
		{
			groupQuestionRoute.GET("", wrapper.Wrap(groupQuestionAPI.GetPaging))
			groupQuestionRoute.GET("/all", wrapper.Wrap(groupQuestionAPI.GetAll))
			groupQuestionRoute.GET("/:id", wrapper.Wrap(groupQuestionAPI.GetByID))
			groupQuestionRoute.POST("", wrapper.Wrap(groupQuestionAPI.Create))
			groupQuestionRoute.PUT("/:id", wrapper.Wrap(groupQuestionAPI.Update))
			groupQuestionRoute.DELETE("/:id", wrapper.Wrap(groupQuestionAPI.Delete))
		}

		// Category
		categoryRoute := api1.Group("/categories", authMiddleware)
		{
			categoryRoute.GET("", wrapper.Wrap(categoryAPI.GetPaging))
			categoryRoute.GET("/all", wrapper.Wrap(categoryAPI.GetAll))
			categoryRoute.GET("/:id", wrapper.Wrap(categoryAPI.GetByID))
			categoryRoute.POST("", wrapper.Wrap(categoryAPI.Create))
			categoryRoute.PUT("/:id", wrapper.Wrap(categoryAPI.Update))
			categoryRoute.DELETE("/:id", wrapper.Wrap(categoryAPI.Delete))
		}

		// Subject
		subjectRoute := api1.Group("/subjects", authMiddleware)
		{
			subjectRoute.POST("", wrapper.Wrap(subjectAPI.Create))
			subjectRoute.PUT("/:id", wrapper.Wrap(subjectAPI.Update))
			subjectRoute.PUT("/move/:id", wrapper.Wrap(subjectAPI.Move))
			subjectRoute.DELETE("/:id/:category_id", wrapper.Wrap(subjectAPI.Delete))
		}

		// Question
		questionRoute := api1.Group("/questions", authMiddleware)
		{
			questionRoute.GET("", wrapper.Wrap(questionAPI.GetPaging))
			questionRoute.GET("/:id", wrapper.Wrap(questionAPI.GetByID))
			questionRoute.POST("", wrapper.Wrap(questionAPI.Create))
			questionRoute.POST("/clones/:id", wrapper.Wrap(questionAPI.Clones))
			questionRoute.PUT("/:id", wrapper.Wrap(questionAPI.Update))
			questionRoute.DELETE("/:id", wrapper.Wrap(questionAPI.Delete))
		}

		// Exams
		examRoute := api1.Group("/exams", authMiddleware)
		{
			examRoute.GET("", wrapper.Wrap(examAPI.GetPaging))
			examRoute.GET("/:id", wrapper.Wrap(examAPI.GetByID))
			examRoute.POST("", wrapper.Wrap(examAPI.Create))
			examRoute.PUT("/:id", wrapper.Wrap(examAPI.Update))
			examRoute.POST("/add/:exam_id/:question_id", wrapper.Wrap(examAPI.AddQuestion))
			examRoute.DELETE("/delete/:exam_id/:question_id", wrapper.Wrap(examAPI.DeleteQuestion))
			examRoute.PUT("/move", wrapper.Wrap(examAPI.MoveQuestion))
			examRoute.DELETE("/:id", wrapper.Wrap(examAPI.Delete))
		}

		// Rooms
		roomRoute := api1.Group("/rooms")
		{
			roomRoute.GET("", authMiddleware, wrapper.Wrap(roomAPI.GetPaging))
			roomRoute.GET("/code", wrapper.Wrap(roomAPI.GetCodeRoom))
			roomRoute.GET("/:id", authMiddleware, wrapper.Wrap(roomAPI.GetByID))
			roomRoute.POST("", authMiddleware, wrapper.Wrap(roomAPI.Create))
			roomRoute.PUT("/:id", authMiddleware, wrapper.Wrap(roomAPI.Update))
			roomRoute.DELETE("/:id", authMiddleware, wrapper.Wrap(roomAPI.Delete))
			// roomRoute.GET("/exam/:id", authMiddleware, wrapper.Wrap(roomAPI.Delete))
			// roomRoute.GET("/result/:id", authMiddleware, wrapper.Wrap(roomAPI.Delete))
		}

		// Candidate
		// candidateRoute := api1.Group("/rooms")
		// {
		// 	candidateRoute.POST("/create", authMiddleware, wrapper.Wrap(roomAPI.GetPaging))
		// 	candidateRoute.POST("/submit", authMiddleware, wrapper.Wrap(roomAPI.Create))
		// }

		// submit
		// get list exam
		//

		return nil
	})

	if err != nil {
		logger.Error(err.Error())
	}

	return err
}
