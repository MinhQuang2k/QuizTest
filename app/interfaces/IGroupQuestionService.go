package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IGroupQuestionService interface {
	GetPaging(c context.Context, req *serializers.GetPagingGroupQuestionReq) ([]*models.GroupQuestion, *paging.Pagination, error)
	GetAll(c context.Context, userID uint) ([]*models.GroupQuestion, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.GroupQuestion, error)
	Create(ctx context.Context, req *serializers.CreateGroupQuestionReq) error
	Update(ctx context.Context, id uint, req *serializers.UpdateGroupQuestionReq) error
	Delete(ctx context.Context, id uint, userID uint) error
}
