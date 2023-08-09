package interfaces

import (
	"context"
	"quiztest/app/serializers"
)

type ISubjectService interface {
	Create(ctx context.Context, req *serializers.CreateSubjectReq) error
	Update(ctx context.Context, id uint, req *serializers.UpdateSubjectReq) error
	Move(ctx context.Context, id uint, req *serializers.MoveSubjectReq) error
	Delete(ctx context.Context, id uint, categoryID uint, userID uint) error
}
