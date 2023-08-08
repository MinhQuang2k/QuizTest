package repositories

import (
	"context"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/config"
	"quiztest/pkg/errors"
)

type UserRepo struct {
	db interfaces.IDatabase
}

func NewUserRepository(db interfaces.IDatabase) interfaces.IUserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.GetInstance().Create(&user).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *UserRepo) Update(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.GetInstance().Save(&user).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *UserRepo) GetByID(ctx context.Context, id uint) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var user models.User
	if err := r.db.GetInstance().Where("id = ? ", id).First(&user).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var user models.User
	if err := r.db.GetInstance().Where("email = ? ", email).First(&user).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &user, nil
}
