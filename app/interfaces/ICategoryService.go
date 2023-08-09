package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type ICategoryService interface {
	GetPaging(c context.Context, req *serializers.GetPagingCategoryReq) ([]*models.Category, *paging.Pagination, error)
	GetAll(c context.Context, userID uint) ([]*models.Category, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Category, error)
	Create(ctx context.Context, req *serializers.CreateCategoryReq) error
	Update(ctx context.Context, id uint, req *serializers.UpdateCategoryReq) error
	Delete(ctx context.Context, id uint, userID uint) error
}
