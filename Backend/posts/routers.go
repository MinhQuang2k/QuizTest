package posts

import (
	"github.com/gin-gonic/gin"
)

func PostRegister(router *gin.RouterGroup) {
	router.POST("/", PostCreate)
	router.PUT("/:slug", PostUpdate)
	router.DELETE("/:slug", PostDelete)
	router.POST("/:slug/comments", PostCommentCreate)
	router.DELETE("/:slug/comments/:id", PostCommentDelete)
}

func PostAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", PostList)
	router.GET("/:slug", PostRetrieve)
	router.GET("/:slug/comments", PostCommentList)
}
