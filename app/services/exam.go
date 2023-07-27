package services

import (
	"context"

	"quiztest/pkg/errors"
	"quiztest/pkg/logger"

	"quiztest/app/models"
	"quiztest/app/repositories"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
	"quiztest/pkg/utils"
)

type IExamService interface {
	GetPaging(c context.Context, req *serializers.GetPagingExamReq) ([]*models.Exam, *paging.Pagination, error)
	GetAll(c context.Context, userID uint) ([]*models.Exam, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Exam, error)
	Create(ctx context.Context, req *serializers.CreateExamReq) (*models.Exam, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateExamReq) (*models.Exam, error)
	AddQuestion(ctx context.Context, id, question_id, userID uint) (*models.Exam, error)
	Delete(ctx context.Context, id uint, userID uint) (*models.Exam, error)
}

type ExamService struct {
	repo         repositories.IExamRepository
	repoQuestion repositories.IQuestionRepository
}

func NewExamService(repo repositories.IExamRepository, repoQuestion repositories.IQuestionRepository) *ExamService {
	return &ExamService{repo: repo, repoQuestion: repoQuestion}
}

func (p *ExamService) GetByID(ctx context.Context, id uint, userID uint) (*models.Exam, error) {
	exam, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return exam, nil
}

func (p *ExamService) GetPaging(ctx context.Context, req *serializers.GetPagingExamReq) ([]*models.Exam, *paging.Pagination, error) {
	exams, pagination, err := p.repo.GetPaging(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return exams, pagination, nil
}
func (p *ExamService) GetAll(ctx context.Context, userID uint) ([]*models.Exam, error) {
	exams, err := p.repo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	return exams, nil
}

func (p *ExamService) Create(ctx context.Context, req *serializers.CreateExamReq) (*models.Exam, error) {
	var exam models.Exam
	utils.Copy(&exam, req)

	err := p.repo.Create(ctx, &exam, req.UserID)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &exam, nil
}

func (p *ExamService) Update(ctx context.Context, id uint, req *serializers.UpdateExamReq) (*models.Exam, error) {
	exam, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Errorf("Update.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	utils.Copy(exam, req)
	err = p.repo.Update(ctx, exam)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return exam, nil
}

func (p *ExamService) Delete(ctx context.Context, id uint, userID uint) (*models.Exam, error) {
	exam, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("Delete.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	err = p.repo.Delete(ctx, exam)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return exam, nil
}

func (p *ExamService) AddQuestion(ctx context.Context, id, question_id, userID uint) (*models.Exam, error) {
	question, err := p.repoQuestion.GetByID(ctx, question_id, userID)
	if err != nil {
		logger.Errorf("AddQuestion.GetQuestion fail, id: %s, error: %s", question_id, err)
		return nil, err
	}

	exam, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("AddQuestion.GetExam fail, id: %s, error: %s", id, err)
		return nil, err
	}

	if utils.FindUint(exam.Questions, question.ID) != 0 {
		logger.Errorf("Question exits fail, id: %s, error: %s", id, err)
		return nil, errors.ErrorDatabaseCreate.Newm("question exist")
	}
	exam.Questions = append(exam.Questions, question.ID)
	err = p.repo.Update(ctx, exam)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return exam, nil
}
