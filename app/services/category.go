package services

import (
	"context"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/logger"
	"quiztest/pkg/paging"
	"quiztest/pkg/utils"
)

type CategoryService struct {
	repo interfaces.ICategoryRepository
}

func NewCategoryService(repo interfaces.ICategoryRepository) interfaces.ICategoryService {
	return &CategoryService{repo: repo}
}

func (p *CategoryService) Create(ctx context.Context, req *serializers.CreateCategoryReq) error {
	var category models.Category
	utils.Copy(&category, req)
	var subjects []*models.Subject
	utils.Copy(&subjects, req.Subjects)
	category.Subjects = subjects

	if err := p.repo.Create(ctx, &category); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (p *CategoryService) GetByID(ctx context.Context, id uint, userID uint) (*models.Category, error) {
	category, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (p *CategoryService) GetPaging(ctx context.Context, req *serializers.GetPagingCategoryReq) ([]*models.Category, *paging.Pagination, error) {
	categories, pagination, err := p.repo.GetPaging(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil
}

func (p *CategoryService) GetAll(ctx context.Context, userID uint) ([]*models.Category, error) {
	categories, err := p.repo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (p *CategoryService) Update(ctx context.Context, id uint, req *serializers.UpdateCategoryReq) error {
	category, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Error(err)
		return err
	}

	utils.Copy(category, req)
	err = p.repo.Update(ctx, category)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (p *CategoryService) Delete(ctx context.Context, id uint, userID uint) error {
	category, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = p.repo.Delete(ctx, category)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
