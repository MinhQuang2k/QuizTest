package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IQuestionService interface {
	GetPaging(c context.Context, req *serializers.GetPagingQuestionReq) ([]*models.Question, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Question, error)
	Clones(ctx context.Context, userID uint, questionClonesID uint) error
	Create(ctx context.Context, req *serializers.CreateQuestionReq) (*models.Question, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateQuestionReq) (*models.Question, error)
	Delete(ctx context.Context, id uint, userID uint) (*models.Question, error)
}
