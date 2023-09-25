package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IQuestionRepository interface {
	Create(ctx context.Context, question *models.Question) error
	Clones(ctx context.Context, userID uint, questionClonesID uint) (*models.Question, error)
	Update(ctx context.Context, question *models.Question, userID uint) error
	GetPaging(ctx context.Context, req *serializers.GetPagingQuestionReq) ([]*models.Question, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Question, error)
	GetByExamID(ctx context.Context, id uint) ([]*models.Question, error)
	Delete(ctx context.Context, question *models.Question) error
	GetTotalScore(ctx context.Context, questions []uint) uint
}
