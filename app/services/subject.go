package services

import (
	"context"

	"quiztest/pkg/logger"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/utils"
)

type SubjectService struct {
	repo interfaces.ISubjectRepository
}

func NewSubjectService(repo interfaces.ISubjectRepository) interfaces.ISubjectService {
	return &SubjectService{repo: repo}
}

func (p *SubjectService) Create(ctx context.Context, req *serializers.CreateSubjectReq) error {
	var subject models.Subject
	utils.Copy(&subject, req)

	err := p.repo.Create(ctx, &subject)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return err
	}

	return nil
}

func (p *SubjectService) Move(ctx context.Context, id uint, req *serializers.MoveSubjectReq) error {
	subject, err := p.repo.GetByID(ctx, id, req.CategoryID)
	if err != nil {
		logger.Errorf("Move.GetByID fail, id: %s, error: %s", id, err)
		return err
	}

	err = p.repo.Move(ctx, req, subject)
	if err != nil {
		logger.Errorf("Move fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}
func (p *SubjectService) Update(ctx context.Context, id uint, req *serializers.UpdateSubjectReq) error {
	subject, err := p.repo.GetByID(ctx, id, req.CategoryID)
	if err != nil {
		logger.Errorf("Update.GetByID fail, id: %s, error: %s", id, err)
		return err
	}

	utils.Copy(subject, req)
	err = p.repo.Update(ctx, subject)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (p *SubjectService) Delete(ctx context.Context, id uint, CategoryID uint, userID uint) error {
	subject, err := p.repo.GetByID(ctx, id, CategoryID)
	if err != nil {
		logger.Errorf("Delete.GetByID fail, id: %s, error: %s", id, err)
		return err
	}

	err = p.repo.Delete(ctx, subject)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}
