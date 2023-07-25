package services

import (
	"context"

	"quiztest/pkg/logger"

	"quiztest/app/models"
	"quiztest/app/repositories"
	"quiztest/app/serializers"
	"quiztest/pkg/utils"
)

type ISubjectService interface {
	Create(ctx context.Context, req *serializers.CreateSubjectReq) (*models.Subject, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateSubjectReq) (*models.Subject, error)
	Delete(ctx context.Context, id uint, userID uint) (*models.Subject, error)
}

type SubjectService struct {
	repo repositories.ISubjectRepository
}

func NewSubjectService(repo repositories.ISubjectRepository) *SubjectService {
	return &SubjectService{repo: repo}
}

func (p *SubjectService) Create(ctx context.Context, req *serializers.CreateSubjectReq) (*models.Subject, error) {
	var subject models.Subject
	utils.Copy(&subject, req)

	err := p.repo.Create(ctx, &subject)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &subject, nil
}

func (p *SubjectService) Update(ctx context.Context, id uint, req *serializers.UpdateSubjectReq) (*models.Subject, error) {
	subject, err := p.repo.GetByID(ctx, id)
	if err != nil {
		logger.Errorf("Update.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	utils.Copy(subject, req)
	err = p.repo.Update(ctx, subject)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return subject, nil
}

func (p *SubjectService) Delete(ctx context.Context, id uint, userID uint) (*models.Subject, error) {
	subject, err := p.repo.GetByID(ctx, id)
	if err != nil {
		logger.Errorf("Delete.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	err = p.repo.Delete(ctx, subject)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return subject, nil
}
