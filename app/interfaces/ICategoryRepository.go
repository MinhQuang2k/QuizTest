package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type ICategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, category *models.Category) error
	GetPaging(ctx context.Context, req *serializers.GetPagingCategoryReq) ([]*models.Category, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Category, error)
	GetAll(ctx context.Context, userID uint) ([]*models.Category, error)
}
