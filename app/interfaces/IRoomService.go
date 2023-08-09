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
	Create(ctx context.Context, req *serializers.CreateRoomReq) error
	Update(ctx context.Context, id uint, req *serializers.UpdateRoomReq) error
	Delete(ctx context.Context, id uint, userID uint) error
}
