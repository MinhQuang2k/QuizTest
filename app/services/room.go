package services

import (
	"context"

	"quiztest/pkg/logger"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
	"quiztest/pkg/utils"

	"github.com/google/uuid"
)

type RoomService struct {
	repo interfaces.IRoomRepository
}

func NewRoomService(repo interfaces.IRoomRepository) interfaces.IRoomService {
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

func (p *RoomService) Create(ctx context.Context, req *serializers.CreateRoomReq) error {
	var room models.Room
	utils.Copy(&room, req)
	room.LinkRoomId = uuid.New().String()
	err := p.repo.Create(ctx, &room)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return err
	}

	return nil
}

func (p *RoomService) Update(ctx context.Context, id uint, req *serializers.UpdateRoomReq) error {
	room, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Errorf("Update.GetByID fail, id: %s, error: %s", id, err)
		return err
	}

	utils.Copy(room, req)
	err = p.repo.Update(ctx, room)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (p *RoomService) Delete(ctx context.Context, id uint, userID uint) error {
	room, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("Delete.GetByID fail, id: %s, error: %s", id, err)
		return err
	}

	err = p.repo.Delete(ctx, room)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}
