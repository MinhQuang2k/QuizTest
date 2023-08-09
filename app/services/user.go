package services

import (
	"context"
	"errors"

	"quiztest/pkg/logger"

	"golang.org/x/crypto/bcrypt"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/jtoken"
	"quiztest/pkg/utils"
)

type UserService struct {
	repo interfaces.IUserRepository
}

func NewUserService(repo interfaces.IUserRepository) interfaces.IUserService {
	return &UserService{repo: repo}
}

func (u *UserService) Login(ctx context.Context, req *serializers.LoginReq) (string, string, error) {
	user, err := u.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		logger.Errorf("Login.GetByEmail fail, email: %s, error: %s", req.Email, err)
		return "", "", err
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return "", "", errors.New("wrong password")
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}
	accessToken := jtoken.GenerateAccessToken(tokenData)
	refreshToken := jtoken.GenerateRefreshToken(tokenData)
	return accessToken, refreshToken, nil
}

func (u *UserService) Register(ctx context.Context, req *serializers.RegisterReq) error {
	var user models.User
	utils.Copy(&user, &req)
	err := u.repo.Create(ctx, &user)
	if err != nil {
		logger.Errorf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return err
	}
	return nil
}

func (u *UserService) GetByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		logger.Errorf("GetByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return user, nil
}

func (u *UserService) RefreshToken(ctx context.Context, userID uint) (string, error) {
	user, err := u.repo.GetByID(ctx, userID)
	if err != nil {
		logger.Errorf("RefreshToken.GetByID fail, id: %s, error: %s", userID, err)
		return "", err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}
	accessToken := jtoken.GenerateAccessToken(tokenData)
	return accessToken, nil
}

func (u *UserService) ChangePassword(ctx context.Context, id uint, req *serializers.ChangePasswordReq) error {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		logger.Errorf("ChangePassword.GetByID fail, id: %s, error: %s", id, err)
		return err
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return errors.New("wrong password")
	}

	user.Password = utils.HashAndSalt([]byte(req.NewPassword))
	err = u.repo.Update(ctx, user)
	if err != nil {
		logger.Errorf("ChangePassword.Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}
