package services

import (
	"context"

	"goshop/pkg/logger"

	"goshop/app/models"
	"goshop/app/repositories"
	"goshop/app/serializers"
	"goshop/pkg/paging"
	"goshop/pkg/utils"
)

type IGroupQuestionService interface {
	ListGroupQuestions(c context.Context, req *serializers.ListGroupQuestionReq) ([]*models.GroupQuestion, *paging.Pagination, error)
	GetGroupQuestionByID(ctx context.Context, id string, userID string) (*models.GroupQuestion, error)
	Create(ctx context.Context, req *serializers.CreateGroupQuestionReq) (*models.GroupQuestion, error)
	Update(ctx context.Context, id string, req *serializers.UpdateGroupQuestionReq) (*models.GroupQuestion, error)
}

type GroupQuestionService struct {
	repo repositories.IGroupQuestionRepository
}

func NewGroupQuestionService(repo repositories.IGroupQuestionRepository) *GroupQuestionService {
	return &GroupQuestionService{repo: repo}
}

func (p *GroupQuestionService) GetGroupQuestionByID(ctx context.Context, id string, userID string) (*models.GroupQuestion, error) {
	groupQuestion, err := p.repo.GetGroupQuestionByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return groupQuestion, nil
}

func (p *GroupQuestionService) ListGroupQuestions(ctx context.Context, req *serializers.ListGroupQuestionReq) ([]*models.GroupQuestion, *paging.Pagination, error) {
	groupQuestions, pagination, err := p.repo.ListGroupQuestions(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return groupQuestions, pagination, nil
}

func (p *GroupQuestionService) Create(ctx context.Context, req *serializers.CreateGroupQuestionReq) (*models.GroupQuestion, error) {
	_, err := p.repo.ExitName(ctx, req.Name, req.UserID)
	if err != nil {
		logger.Errorf("Create exits name, name: %s, error: %s", req.Name, err)
		return nil, err
	}

	var groupQuestion models.GroupQuestion
	utils.Copy(&groupQuestion, req)

	err = p.repo.Create(ctx, &groupQuestion)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &groupQuestion, nil
}

func (p *GroupQuestionService) Update(ctx context.Context, id string, req *serializers.UpdateGroupQuestionReq) (*models.GroupQuestion, error) {
	groupQuestion, err := p.repo.GetGroupQuestionByID(ctx, id, req.UserID)
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
