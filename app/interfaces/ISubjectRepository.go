package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
)

type ISubjectRepository interface {
	Create(ctx context.Context, subject *models.Subject) error
	Update(ctx context.Context, subject *models.Subject) error
	Move(ctx context.Context, req *serializers.MoveSubjectReq, subject *models.Subject) error
	Delete(ctx context.Context, subject *models.Subject) error
	GetByID(ctx context.Context, id uint, categoryID uint) (*models.Subject, error)
}
