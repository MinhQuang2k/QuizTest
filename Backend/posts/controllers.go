package posts

import (
	"errors"
	"net/http"
	"strconv"

	"blog.com/common"
	"blog.com/models"
	"github.com/gin-gonic/gin"
)

func PostCommentDelete(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("Invalid id")))
		return
	}
	err = models.DeleteComment([]uint{id})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("Invalid id")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": "Delete success"})
}

func PostCommentList(c *gin.Context) {
	slug := c.Param("slug")
	postModel, err := models.FindPost(&models.PostModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("Invalid slug")))
		return
	}
	err = postModel.GetComments()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("Database error")))
		return
	}
	serializer := CommentsSerializer{c, postModel.Comments}
	c.JSON(http.StatusOK, gin.H{"comments": serializer.Response()})
}

func PostCommentCreate(c *gin.Context) {
	slug := c.Param("slug")
	postModel, err := models.FindPost(&models.PostModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("Invalid slug")))
		return
	}
	commentModelValidator := NewCommentModelValidator()
	if err := commentModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	commentModelValidator.commentModel.Post = postModel

	if err := models.SaveOne(&commentModelValidator.commentModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := CommentSerializer{c, commentModelValidator.commentModel}
	c.JSON(http.StatusCreated, gin.H{"comment": serializer.Response()})
}

func PostRetrieve(c *gin.Context) {
	slug := c.Param("slug")
	postModel, err := models.FindPost(&models.PostModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("posts", errors.New("Invalid slug")))
		return
	}
	serializer := PostSerializer{c, postModel}
	c.JSON(http.StatusOK, gin.H{"post": serializer.Response()})
}

func PostCreate(c *gin.Context) {
	postModelValidator := NewPostModelValidator()
	if err := postModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := models.SaveOne(&postModelValidator.postModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := PostSerializer{c, postModelValidator.postModel}
	c.JSON(http.StatusCreated, gin.H{"post": serializer.Response()})
}

func PostUpdate(c *gin.Context) {
	slug := c.Param("slug")
	postModel, err := models.FindPost(&models.PostModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("posts", errors.New("Invalid slug")))
		return
	}
	postModelValidator := NewPostModelValidatorFillWith(postModel)
	if err := postModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	postModelValidator.postModel.ID = postModel.ID
	if err := postModel.Update(postModelValidator.postModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := PostSerializer{c, postModel}
	c.JSON(http.StatusOK, gin.H{"post": serializer.Response()})
}

func PostDelete(c *gin.Context) {
	slug := c.Param("slug")
	err := models.DeletePost(&models.PostModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("posts", errors.New("Invalid slug")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": "Delete success"})
}

func PostList(c *gin.Context) {
	author := c.Query("author")
	limit := c.Query("limit")
	offset := c.Query("offset")
	postModels, modelCount, err := models.FindPostPaging(author, limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("posts", errors.New("Invalid param")))
		return
	}
	serializer := PostsSerializer{c, postModels}
	c.JSON(http.StatusOK, gin.H{"posts": serializer.Response(), "postsCount": modelCount})
}
