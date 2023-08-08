package repositories

import (
	"context"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/config"
	"quiztest/pkg/errors"
	"quiztest/pkg/paging"
)

type RoomRepo struct {
	db interfaces.IDatabase
}

func NewRoomRepository(db interfaces.IDatabase) interfaces.IRoomRepository {
	return &RoomRepo{db: db}
}

func (r *RoomRepo) GetPaging(ctx context.Context, req *serializers.GetPagingRoomReq) ([]*models.Room, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	var total int64
	query := r.db.GetInstance().
		Joins("JOIN exams ON exams.id = rooms.exam_id").
		Joins("JOIN subjects ON subjects.id = exams.subject_id").
		Joins("JOIN categories ON categories.id = subjects.category_id").
		Where("categories.user_id = ?", req.UserID)

	order := "created_at"
	if req.Name != "" {
		query = query.
			Where("name LIKE ?", "%"+req.Name+"%").
			Where("user_id = ?", req.UserID)
	}
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var rooms []*models.Room
	if err := query.
		Limit(int(pagination.Limit)).
		Offset(int(pagination.Skip)).
		Order(order).
		Find(&rooms).Count(&total).Error; err != nil {
		return nil, nil, nil
	}

	pagination.Total = total

	return rooms, pagination, nil
}

func (r *RoomRepo) GetByID(ctx context.Context, id uint, userID uint) (*models.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var room models.Room
	if err := r.db.GetInstance().Joins("JOIN exams ON exams.id = rooms.exam_id").
		Joins("JOIN subjects ON subjects.id = exams.subject_id").
		Joins("JOIN categories ON categories.id = subjects.category_id").
		Where("categories.user_id = ?", userID).
		Where("rooms.id = ?", id).First(&room).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &room, nil
}

func (r *RoomRepo) Create(ctx context.Context, room *models.Room) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.GetInstance().Create(&room).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *RoomRepo) Update(ctx context.Context, room *models.Room) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.GetInstance().Save(&room).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *RoomRepo) Delete(ctx context.Context, room *models.Room) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	rowsAffected := r.db.GetInstance().Delete(&room).RowsAffected

	if rowsAffected == 0 {
		return errors.ErrorNotFound.New()
	}

	return nil
}
