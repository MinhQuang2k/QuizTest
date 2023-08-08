package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IRoomService interface {
	GetPaging(c context.Context, req *serializers.GetPagingRoomReq) ([]*models.Room, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Room, error)
	Create(ctx context.Context, req *serializers.CreateRoomReq) (*models.Room, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateRoomReq) (*models.Room, error)
	Delete(ctx context.Context, id uint, userID uint) (*models.Room, error)
}
