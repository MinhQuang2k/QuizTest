package repositories

import (
	"context"

	"gorm.io/gorm"

	"quiztest/app/dbs"
	"quiztest/app/models"
	"quiztest/config"
	"quiztest/pkg/errors"
)

type ISubjectRepository interface {
	Create(ctx context.Context, subject *models.Subject) error
	CreateMany(ctx context.Context, subjects []*models.Subject, category *models.Category) error
	Update(ctx context.Context, subject *models.Subject) error
	Delete(ctx context.Context, subject *models.Subject) error
	GetByCategoryID(ctx context.Context, categoryID string) (*models.Subject, error)
	GetByID(ctx context.Context, id string) (*models.Subject, error)
	GetAll(ctx context.Context, categoryID string) ([]*models.Subject, error)
}

type SubjectRepo struct {
	db *gorm.DB
}

func NewSubjectRepository() *SubjectRepo {
	return &SubjectRepo{db: dbs.Database}
}

func (r *SubjectRepo) Create(ctx context.Context, subject *models.Subject) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Where("name = ?", subject.Name).Where("category_id = ?", subject.CategoryID).First(&subject).Error; err == nil {
		return errors.ErrorExistName.New()
	}

	if err := r.db.Create(&subject).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *SubjectRepo) CreateMany(ctx context.Context, subjects []*models.Subject, category *models.Category) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	for _, subject := range subjects {
		subject.CategoryID = category.ID
	}

	if err := r.db.Create(&subjects).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *SubjectRepo) Update(ctx context.Context, subject *models.Subject) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Where("name = ?", subject.Name).Where("category_id = ?", subject.CategoryID).First(&subject).Error; err == nil {
		return errors.ErrorExistName.New()
	}

	if err := r.db.Save(&subject).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *SubjectRepo) Delete(ctx context.Context, subject *models.Subject) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	rowsAffected := r.db.Delete(&subject).RowsAffected

	if rowsAffected == 0 {
		return errors.ErrorNotFound.New()
	}

	return nil
}

func (r *SubjectRepo) GetAll(ctx context.Context, categoryID string) ([]*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var subjects []*models.Subject
	if err := r.db.Where("categoryID = ?", categoryID).Find(&subjects).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return subjects, nil
}

func (r *SubjectRepo) GetByCategoryID(ctx context.Context, categoryID string) (*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var subject models.Subject
	if err := r.db.Where("category_id = ?", categoryID).First(&subject).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &subject, nil
}

func (r *SubjectRepo) GetByID(ctx context.Context, id string) (*models.Subject, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var subject models.Subject
	if err := r.db.Where("id = ?", id).First(&subject).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &subject, nil
}
