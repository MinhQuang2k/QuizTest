package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IGroupQuestionRepository interface {
	Create(ctx context.Context, groupQuestion *models.GroupQuestion) error
	Update(ctx context.Context, groupQuestion *models.GroupQuestion) error
	GetPaging(ctx context.Context, req *serializers.GetPagingGroupQuestionReq) ([]*models.GroupQuestion, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.GroupQuestion, error)
	GetAll(ctx context.Context, userID uint) ([]*models.GroupQuestion, error)
	Delete(ctx context.Context, groupQuestion *models.GroupQuestion) error
}
