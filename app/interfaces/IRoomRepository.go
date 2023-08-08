package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IRoomRepository interface {
	Create(ctx context.Context, room *models.Room) error
	Update(ctx context.Context, room *models.Room) error
	GetPaging(ctx context.Context, req *serializers.GetPagingRoomReq) ([]*models.Room, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Room, error)
	Delete(ctx context.Context, room *models.Room) error
}
