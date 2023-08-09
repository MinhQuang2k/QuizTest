package repositories

import (
	"context"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/config"
	"quiztest/pkg/errors"
)

type SubjectRepo struct {
	db interfaces.IDatabase
}

func NewSubjectRepository(db interfaces.IDatabase) interfaces.ISubjectRepository {
	return &SubjectRepo{db: db}
}

func (r *SubjectRepo) Create(ctx context.Context, subject *models.Subject) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if rowsAffected := r.db.GetInstance().Where("name = ?", subject.Name).
		Where("category_id = ?", subject.CategoryID).
		First(&subject).RowsAffected; rowsAffected != 0 {
		return errors.ErrorExistName.New()
	}

	if err := r.db.GetInstance().Create(&subject).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *SubjectRepo) Update(ctx context.Context, subject *models.Subject) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if rowsAffected := r.db.GetInstance().Where("name = ?", subject.Name).
		Where("category_id = ?", subject.CategoryID).
		First(&subject).RowsAffected; rowsAffected != 0 {
		return errors.ErrorExistName.New()
	}

	if err := r.db.GetInstance().Save(&subject).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}
func (r *SubjectRepo) Move(ctx context.Context, req *serializers.MoveSubjectReq, subject *models.Subject) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	var newSubject models.Subject
	if rowsAffected := r.db.GetInstance().Where("category_id = ?", req.NewCategoryID).
		Where("name = ?", subject.Name).
		First(&newSubject).RowsAffected; rowsAffected != 0 {
		return errors.ErrorExistName.New()
	}

	subject.CategoryID = req.NewCategoryID

	if err := r.db.GetInstance().Save(&subject).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *SubjectRepo) Delete(ctx context.Context, subject *models.Subject) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.GetInstance().Delete(&subject).Error; err != nil {
		return errors.ErrorDatabaseDelete.New()
	}

	return nil
}

func (r *SubjectRepo) GetAll(ctx context.Context, categoryID uint) ([]*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var subjects []*models.Subject
	if err := r.db.GetInstance().Where("categoryID = ?", categoryID).Find(&subjects).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return subjects, nil
}

func (r *SubjectRepo) GetByID(ctx context.Context, id uint, categoryID uint) (*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var subject models.Subject
	if err := r.db.GetInstance().Where("id = ?", id).
		Where("category_id = ?", categoryID).
		First(&subject).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &subject, nil
}
