package services

import (
	"context"

	"quiztest/pkg/logger"

	"quiztest/app/models"
	"quiztest/app/repositories"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
	"quiztest/pkg/utils"
)

type IRoomService interface {
	GetPaging(c context.Context, req *serializers.GetPagingRoomReq) ([]*models.Room, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Room, error)
	Create(ctx context.Context, req *serializers.CreateRoomReq) (*models.Room, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateRoomReq) (*models.Room, error)
	Delete(ctx context.Context, id uint, userID uint) (*models.Room, error)
}

type RoomService struct {
	repo repositories.IRoomRepository
}

func NewRoomService(repo repositories.IRoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (p *RoomService) GetByID(ctx context.Context, id uint, userID uint) (*models.Room, error) {
	room, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (p *RoomService) GetPaging(ctx context.Context, req *serializers.GetPagingRoomReq) ([]*models.Room, *paging.Pagination, error) {
	rooms, pagination, err := p.repo.GetPaging(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return rooms, pagination, nil
}

func (p *RoomService) Create(ctx context.Context, req *serializers.CreateRoomReq) (*models.Room, error) {
	var room models.Room
	utils.Copy(&room, req)

	err := p.repo.Create(ctx, &room)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &room, nil
}

func (p *RoomService) Update(ctx context.Context, id uint, req *serializers.UpdateRoomReq) (*models.Room, error) {
	room, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Errorf("Update.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	utils.Copy(room, req)
	err = p.repo.Update(ctx, room)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return room, nil
}

func (p *RoomService) Delete(ctx context.Context, id uint, userID uint) (*models.Room, error) {
	room, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("Delete.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	err = p.repo.Delete(ctx, room)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return room, nil
}
