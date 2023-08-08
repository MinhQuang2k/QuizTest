package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
)

type ISubjectService interface {
	Create(ctx context.Context, req *serializers.CreateSubjectReq) (*models.Subject, error)
	Update(ctx context.Context, id uint, req *serializers.UpdateSubjectReq) (*models.Subject, error)
	Move(ctx context.Context, id uint, req *serializers.MoveSubjectReq) (*models.Subject, error)
	Delete(ctx context.Context, id uint, categoryID uint, userID uint) (*models.Subject, error)
}
