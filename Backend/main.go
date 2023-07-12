package main

import (
	"blog.com/common"
	"blog.com/models"
	"blog.com/posts"
	"blog.com/users"
	"github.com/gin-gonic/gin"
)

func main() {

	db := common.Init()
	models.Migrate(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	posts.PostAnonymousRegister(v1.Group("/posts"))
	v1.Use(users.AuthMiddleware(true))
	posts.PostRegister(v1.Group("/posts"))
	users.UserRegister(v1.Group("/user"))

	r.Run()
}
