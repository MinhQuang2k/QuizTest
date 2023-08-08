package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IExamRepository interface {
	Create(ctx context.Context, exam *models.Exam, userID uint) error
	Update(ctx context.Context, exam *models.Exam) error
	UpdateExamQuestion(ctx context.Context, examQuestion *models.ExamQuestion) error
	GetPaging(ctx context.Context, req *serializers.GetPagingExamReq) ([]*models.Exam, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Exam, error)
	GetExamQuestionByID(ctx context.Context, examID uint, questionID uint) (*models.ExamQuestion, error)
	GetAll(ctx context.Context, userID uint) ([]*models.Exam, error)
	Delete(ctx context.Context, exam *models.Exam) error
	DeleteExamQuestion(ctx context.Context, examQuestion *models.ExamQuestion) error
}
