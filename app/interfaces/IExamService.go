package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
)

type IExamService interface {
	GetPaging(c context.Context, req *serializers.GetPagingExamReq) ([]*serializers.Exam, *paging.Pagination, error)
	GetAll(c context.Context, userID uint) ([]*models.Exam, error)
	GetByID(ctx context.Context, id uint, userID uint) (*serializers.Exam, []*models.Question, error)
	Create(ctx context.Context, req *serializers.CreateExamReq) error
	Update(ctx context.Context, id uint, req *serializers.UpdateExamReq) error
	AddQuestion(ctx context.Context, id, question_id, userID uint) error
	Delete(ctx context.Context, id uint, userID uint) error
	DeleteQuestion(ctx context.Context, id, question_id, userID uint) error
	MoveQuestion(ctx context.Context, req *serializers.MoveExamReq) error
}
