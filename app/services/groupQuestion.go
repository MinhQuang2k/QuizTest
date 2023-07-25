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

type IGroupQuestionService interface {
	GetPaging(c context.Context, req *serializers.GetPagingGroupQuestionReq) ([]*models.GroupQuestion, *paging.Pagination, error)
	GetAll(c context.Context, userID uint) ([]*models.GroupQuestion, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.GroupQuestion, error)
	Create(ctx context.Context, req *serializers.CreateGroupQuestionReq) (*models.GroupQuestion, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateGroupQuestionReq) (*models.GroupQuestion, error)
	Delete(ctx context.Context, id uint, userID uint) (*models.GroupQuestion, error)
}

type GroupQuestionService struct {
	repo repositories.IGroupQuestionRepository
}

func NewGroupQuestionService(repo repositories.IGroupQuestionRepository) *GroupQuestionService {
	return &GroupQuestionService{repo: repo}
}

func (p *GroupQuestionService) GetByID(ctx context.Context, id uint, userID uint) (*models.GroupQuestion, error) {
	groupQuestion, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return groupQuestion, nil
}

func (p *GroupQuestionService) GetPaging(ctx context.Context, req *serializers.GetPagingGroupQuestionReq) ([]*models.GroupQuestion, *paging.Pagination, error) {
	groupQuestions, pagination, err := p.repo.GetPaging(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return groupQuestions, pagination, nil
}
func (p *GroupQuestionService) GetAll(ctx context.Context, userID uint) ([]*models.GroupQuestion, error) {
	groupQuestions, err := p.repo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	return groupQuestions, nil
}

func (p *GroupQuestionService) Create(ctx context.Context, req *serializers.CreateGroupQuestionReq) (*models.GroupQuestion, error) {
	var groupQuestion models.GroupQuestion
	utils.Copy(&groupQuestion, req)

	err := p.repo.Create(ctx, &groupQuestion)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &groupQuestion, nil
}

func (p *GroupQuestionService) Update(ctx context.Context, id uint, req *serializers.UpdateGroupQuestionReq) (*models.GroupQuestion, error) {
	groupQuestion, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Errorf("Update.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	utils.Copy(groupQuestion, req)
	err = p.repo.Update(ctx, groupQuestion)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return groupQuestion, nil
}

func (p *GroupQuestionService) Delete(ctx context.Context, id uint, userID uint) (*models.GroupQuestion, error) {
	groupQuestion, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("Delete.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	err = p.repo.Delete(ctx, groupQuestion)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return groupQuestion, nil
}
