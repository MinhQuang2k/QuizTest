package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IExamService interface {
	GetPaging(c context.Context, req *serializers.GetPagingExamReq) ([]*models.Exam, *paging.Pagination, error)
	GetAll(c context.Context, userID uint) ([]*models.Exam, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Exam, []*models.Question, error)
	Create(ctx context.Context, req *serializers.CreateExamReq) (*models.Exam, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateExamReq) (*models.Exam, error)
	AddQuestion(ctx context.Context, id, question_id, userID uint) (*models.Exam, error)
	Delete(ctx context.Context, id uint, userID uint) (*models.Exam, error)
	DeleteQuestion(ctx context.Context, id, question_id, userID uint) error
	MoveQuestion(ctx context.Context, req *serializers.MoveExamReq) error
}
