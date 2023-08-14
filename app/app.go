package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/gin-contrib/cors"
	"quiztest/app/api"
	"quiztest/app/cache"
	"quiztest/app/dbs"
	"quiztest/app/repositories"
	"quiztest/app/router"
	"quiztest/app/services"
	"quiztest/pkg/logger"
)

// BuildContainer build dig container
func BuildContainer() *dig.Container {
	container := dig.New()

	// auth, err := InitAuth()
	// _ = container.Provide(func() jwt.IJWTAuth {
	// 	return auth
	// })

	// Inject database
	err := dbs.Inject(container)
	if err != nil {
		logger.Error("Failed to inject database", err)
	}

	// Inject cache
	err = cache.Inject(container)
	if err != nil {
		logger.Error("Failed to inject cache", err)
	}

	// Inject repositories
	err = repositories.Inject(container)
	if err != nil {
		logger.Error("Failed to inject repositories", err)
	}

	// Inject services
	err = services.Inject(container)
	if err != nil {
		logger.Error("Failed to inject services", err)
	}

	// Inject APIs
	err = api.Inject(container)
	if err != nil {
		logger.Error("Failed to inject APIs", err)
	}

	return container
}

// InitGinEngine initial new gin engine
func InitGinEngine(container *dig.Container) *gin.Engine {
	app := gin.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.Docs(app)
	err := router.RegisterAPI(app, container)
	if err != nil {
		return nil
	}

	return app
}
