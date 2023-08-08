package interfaces

import (
	"context"
	"quiztest/app/models"
	"quiztest/app/serializers"
)

// IUserService interface
type IUserService interface {
	Login(ctx context.Context, req *serializers.LoginReq) (*models.User, string, string, error)
	Register(ctx context.Context, req *serializers.RegisterReq) (*models.User, error)
	GetByID(ctx context.Context, id uint) (*models.User, error)
	RefreshToken(ctx context.Context, userID uint) (string, error)
	ChangePassword(ctx context.Context, id uint, req *serializers.ChangePasswordReq) error
}
