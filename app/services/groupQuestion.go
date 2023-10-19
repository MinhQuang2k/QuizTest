package services

import (
	"context"

	"quiztest/pkg/logger"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
	"quiztest/pkg/utils"
)

type GroupQuestionService struct {
	repo interfaces.IGroupQuestionRepository
}

func NewGroupQuestionService(repo interfaces.IGroupQuestionRepository) interfaces.IGroupQuestionService {
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

func (p *GroupQuestionService) Create(ctx context.Context, req *serializers.CreateGroupQuestionReq) error {
	var groupQuestion models.GroupQuestion
	utils.Copy(&groupQuestion, req)

	err := p.repo.Create(ctx, &groupQuestion)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (p *GroupQuestionService) Update(ctx context.Context, id uint, req *serializers.UpdateGroupQuestionReq) error {
	groupQuestion, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Error(err)
		return err
	}

	utils.Copy(groupQuestion, req)
	err = p.repo.Update(ctx, groupQuestion)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (p *GroupQuestionService) Delete(ctx context.Context, id uint, userID uint) error {
	groupQuestion, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = p.repo.Delete(ctx, groupQuestion)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
