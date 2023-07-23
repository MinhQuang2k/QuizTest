package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"goshop/pkg/logger"
	"goshop/pkg/validation"

	"goshop/app"
	"goshop/app/api"
	"goshop/app/dbs"
	"goshop/app/repositories"
	"goshop/app/services"
	"goshop/config"
)

func main() {
	cfg := config.GetConfig()
	logger.Initialize(cfg.Environment)

	dbs.Init()

	validator := validation.New()

	userRepo := repositories.NewUserRepository()
	groupQuestionRepo := repositories.NewGroupQuestionRepository()

	userSvc := services.NewUserService(userRepo)
	groupQuestionSvc := services.NewGroupQuestionService(groupQuestionRepo)

	userAPI := api.NewUserAPI(validator, userSvc)
	groupQuestionAPI := api.NewGroupQuestionAPI(validator, groupQuestionSvc)

	engine := app.InitGinEngine(userAPI, groupQuestionAPI)

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
