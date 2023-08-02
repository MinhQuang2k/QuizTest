package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"quiztest/pkg/logger"
	"quiztest/pkg/validation"

	"quiztest/app"
	"quiztest/app/api"
	"quiztest/app/dbs"
	"quiztest/app/repositories"
	"quiztest/app/services"
	"quiztest/config"
)

func main() {
	cfg := config.GetConfig()
	logger.Initialize(cfg.Environment)

	dbs.Init()

	validator := validation.New()

	userRepo := repositories.NewUserRepository()
	groupQuestionRepo := repositories.NewGroupQuestionRepository()
	subjectRepo := repositories.NewSubjectRepository()
	categoryRepo := repositories.NewCategoryRepository()
	questionRepo := repositories.NewQuestionRepository()
	examRepo := repositories.NewExamRepository()
	roomRepo := repositories.NewRoomRepository()

	userSvc := services.NewUserService(userRepo)
	groupQuestionSvc := services.NewGroupQuestionService(groupQuestionRepo)
	subjectSvc := services.NewSubjectService(subjectRepo)
	categorySvc := services.NewCategoryService(categoryRepo)
	questionSvc := services.NewQuestionService(questionRepo)
	examSvc := services.NewExamService(examRepo, questionRepo)
	roomSvc := services.NewRoomService(roomRepo)

	userAPI := api.NewUserAPI(validator, userSvc)
	groupQuestionAPI := api.NewGroupQuestionAPI(validator, groupQuestionSvc)
	subjectAPI := api.NewSubjectAPI(validator, subjectSvc)
	categoryAPI := api.NewCategoryAPI(validator, categorySvc)
	questionAPI := api.NewQuestionAPI(validator, questionSvc)
	examAPI := api.NewExamAPI(validator, examSvc)
	roomAPI := api.NewRoomAPI(validator, roomSvc)

	engine := app.InitGinEngine(userAPI, groupQuestionAPI, categoryAPI, subjectAPI, questionAPI, examAPI, roomAPI)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: engine,
	}

	go func() {
		logger.Infof("Listen at: %d\n", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logger.Info("Timeout of 5 seconds.")
	}
	logger.Info("Server exiting")
}
