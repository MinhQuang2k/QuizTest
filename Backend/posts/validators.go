package posts

import (
	"blog.com/common"
	"blog.com/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type PostModelValidator struct {
	Post struct {
		Title       string `form:"title" json:"title" binding:"min=4"`
		Description string `form:"description" json:"description" binding:"max=2048"`
		Body        string `form:"body" json:"body" binding:"max=2048"`
	} `json:"post"`
	postModel models.PostModel `json:"-"`
}

func NewPostModelValidator() PostModelValidator {
	return PostModelValidator{}
}

func NewPostModelValidatorFillWith(postModel models.PostModel) PostModelValidator {
	postModelValidator := NewPostModelValidator()
	postModelValidator.Post.Title = postModel.Title
	postModelValidator.Post.Description = postModel.Description
	postModelValidator.Post.Body = postModel.Body
	return postModelValidator
}

func (s *PostModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.postModel.Slug = slug.Make(s.Post.Title)
	s.postModel.Title = s.Post.Title
	s.postModel.Description = s.Post.Description
	s.postModel.Body = s.Post.Body
	s.postModel.User = myUserModel
	return nil
}

type CommentModelValidator struct {
	Comment struct {
		Body string `form:"body" json:"body" binding:"max=2048"`
	} `json:"comment"`
	commentModel models.CommentModel `json:"-"`
}

func NewCommentModelValidator() CommentModelValidator {
	return CommentModelValidator{}
}

func (s *CommentModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.commentModel.Body = s.Comment.Body
	s.commentModel.User = myUserModel
	return nil
}
