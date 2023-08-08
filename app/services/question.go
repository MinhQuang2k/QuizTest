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

type QuestionService struct {
	repo interfaces.IQuestionRepository
}

func NewQuestionService(repo interfaces.IQuestionRepository) interfaces.IQuestionService {
	return &QuestionService{repo: repo}
}

func (p *QuestionService) GetByID(ctx context.Context, id uint, userID uint) (*models.Question, error) {
	question, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (p *QuestionService) GetPaging(ctx context.Context, req *serializers.GetPagingQuestionReq) ([]*models.Question, *paging.Pagination, error) {
	questions, pagination, err := p.repo.GetPaging(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return questions, pagination, nil
}

func (p *QuestionService) Create(ctx context.Context, req *serializers.CreateQuestionReq) (*models.Question, error) {
	var question models.Question
	utils.Copy(&question, req)

	err := p.repo.Create(ctx, &question)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &question, nil
}

func (p *QuestionService) Clones(ctx context.Context, userID uint, questionClonesID uint) error {
	err := p.repo.Clones(ctx, userID, questionClonesID)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return err
	}

	return nil
}

func (p *QuestionService) Update(ctx context.Context, id uint, req *serializers.UpdateQuestionReq) (*models.Question, error) {
	question, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Errorf("Update.GetByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	utils.Copy(question, req)
	err = p.repo.Update(ctx, question, req.UserID)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return question, nil
}

func (p *QuestionService) Delete(ctx context.Context, id uint, userID uint) (*models.Question, error) {
	question, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("Delete.GetByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	err = p.repo.Delete(ctx, question)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return question, nil
}
