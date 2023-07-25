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

type ICategoryService interface {
	GetPaging(c context.Context, req *serializers.GetPagingCategoryReq) ([]*models.Category, *paging.Pagination, error)
	GetAll(c context.Context, userID uint) ([]*models.Category, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Category, error)
	Create(ctx context.Context, req *serializers.CreateCategoryReq) (*models.Category, []*models.Subject, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateCategoryReq) (*models.Category, error)
	Delete(ctx context.Context, id uint, userID uint) (*models.Category, error)
}

type CategoryService struct {
	repo    repositories.ICategoryRepository
	repoSub repositories.ISubjectRepository
}

func NewCategoryService(repo repositories.ICategoryRepository, repoSub repositories.ISubjectRepository) *CategoryService {
	return &CategoryService{repo: repo, repoSub: repoSub}
}

func (p *CategoryService) Create(ctx context.Context, req *serializers.CreateCategoryReq) (*models.Category, []*models.Subject, error) {
	var category models.Category
	utils.Copy(&category, req)
	var subjects []*models.Subject
	utils.Copy(&subjects, req.Subjects)

	if err := p.repo.Create(ctx, &category); err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, nil, err
	}

	if err := p.repoSub.CreateMany(ctx, subjects, &category); err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, nil, err
	}

	return &category, subjects, nil
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

func (p *CategoryService) Update(ctx context.Context, id uint, req *serializers.UpdateCategoryReq) (*models.Category, error) {
	category, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Errorf("Update.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	utils.Copy(category, req)
	err = p.repo.Update(ctx, category)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return category, nil
}

func (p *CategoryService) Delete(ctx context.Context, id uint, userID uint) (*models.Category, error) {
	category, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("Delete.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	err = p.repo.Delete(ctx, category)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return category, nil
}
