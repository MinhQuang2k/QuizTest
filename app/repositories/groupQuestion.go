package repositories

import (
	"context"

	"gorm.io/gorm"

	"goshop/app/dbs"
	"goshop/app/models"
	"goshop/app/serializers"
	"goshop/config"
	"goshop/pkg/errors"
	"goshop/pkg/paging"
)

type IGroupQuestionRepository interface {
	Create(ctx context.Context, groupQuestion *models.GroupQuestion) error
	Update(ctx context.Context, groupQuestion *models.GroupQuestion) error
	ListGroupQuestions(ctx context.Context, req *serializers.ListGroupQuestionReq) ([]*models.GroupQuestion, *paging.Pagination, error)
	GetGroupQuestionByID(ctx context.Context, id string, userID string) (*models.GroupQuestion, error)
	ExitName(ctx context.Context, name string, userID string) (*models.GroupQuestion, error)
}

type GroupQuestionRepo struct {
	db *gorm.DB
}

func NewGroupQuestionRepository() *GroupQuestionRepo {
	return &GroupQuestionRepo{db: dbs.Database}
}

func (r *GroupQuestionRepo) ListGroupQuestions(ctx context.Context, req *serializers.ListGroupQuestionReq) ([]*models.GroupQuestion, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	query := r.db
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
	var total int64
	if err := query.Model(&models.GroupQuestion{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var groupQuestions []*models.GroupQuestion
	if err := query.
		Limit(int(pagination.Limit)).
		Offset(int(pagination.Skip)).
		Order(order).
		Find(&groupQuestions).Error; err != nil {
		return nil, nil, nil
	}

	return groupQuestions, pagination, nil
}

func (r *GroupQuestionRepo) GetGroupQuestionByID(ctx context.Context, id string, userID string) (*models.GroupQuestion, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var groupQuestion models.GroupQuestion
	if err := r.db.Where("id = ?", id).Where("user_id = ?", userID).First(&groupQuestion).Error; err != nil {
		return nil, err
	}

	return &groupQuestion, nil
}

func (r *GroupQuestionRepo) Create(ctx context.Context, groupQuestion *models.GroupQuestion) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Create(&groupQuestion).Error; err != nil {
		return err
	}

	return nil
}

func (r *GroupQuestionRepo) Update(ctx context.Context, groupQuestion *models.GroupQuestion) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Where("name = ?", groupQuestion.Name).Where("user_id = ?", groupQuestion.UserID).First(&groupQuestion).Error; err == nil {
		return errors.ErrorExistName.Newm(err.Error())
	}

	if err := r.db.Save(&groupQuestion).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}
