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

type CategoryRepo struct {
	db interfaces.IDatabase
}

func NewCategoryRepository(db interfaces.IDatabase) interfaces.ICategoryRepository {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) Create(ctx context.Context, category *models.Category) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if rowsAffected := r.db.GetInstance().
		Where("name = ?", category.Name).
		Where("user_id = ?", category.UserID).
		First(&category).RowsAffected; rowsAffected != 0 {
		return errors.ErrorExistName.New()
	}

	if err := r.db.GetInstance().Create(&category).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *CategoryRepo) GetPaging(ctx context.Context, req *serializers.GetPagingCategoryReq) ([]*models.Category, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	query := r.db.GetInstance()
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
	if err := query.Model(&models.Category{}).Count(&total).Error; err != nil {
		return nil, nil, errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var categories []*models.Category
	if err := query.
		Preload("Subjects").
		Limit(int(pagination.Limit)).
		Offset(int(pagination.Skip)).
		Order(order).
		Find(&categories).Error; err != nil {
		return nil, nil, nil
	}

	return categories, pagination, nil
}

func (r *CategoryRepo) GetAll(ctx context.Context, userID uint) ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var categories []*models.Category
	if err := r.db.GetInstance().Where("user_id = ?", userID).
		Preload("Subjects").
		Find(&categories).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return categories, nil
}

func (r *CategoryRepo) GetByID(ctx context.Context, id uint, userID uint) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var category models.Category
	if err := r.db.GetInstance().
		Where("id = ?", id).
		Where("user_id = ?", userID).
		First(&category).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &category, nil
}

func (r *CategoryRepo) Update(ctx context.Context, category *models.Category) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if rowsAffected := r.db.GetInstance().Where("name = ?", category.Name).
		Where("user_id = ?", category.UserID).
		First(&category).RowsAffected; rowsAffected != 0 {
		return errors.ErrorExistName.New()
	}

	if err := r.db.GetInstance().Save(&category).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *CategoryRepo) Delete(ctx context.Context, category *models.Category) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	rowsAffected := r.db.GetInstance().Delete(&category).RowsAffected

	if rowsAffected == 0 {
		return errors.ErrorDatabaseDelete.New()
	}

	return nil
}
