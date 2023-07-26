package repositories

import (
	"context"

	"gorm.io/gorm"

	"quiztest/app/dbs"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/config"
	"quiztest/pkg/errors"
	"quiztest/pkg/paging"
)

type IExamRepository interface {
	Create(ctx context.Context, exam *models.Exam) error
	Update(ctx context.Context, exam *models.Exam) error
	GetPaging(ctx context.Context, req *serializers.GetPagingExamReq) ([]*models.Exam, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Exam, error)
	GetAll(ctx context.Context, userID uint) ([]*models.Exam, error)
	Delete(ctx context.Context, exam *models.Exam) error
}

type ExamRepo struct {
	db *gorm.DB
}

func NewExamRepository() *ExamRepo {
	return &ExamRepo{db: dbs.Database}
}

func (r *ExamRepo) GetPaging(ctx context.Context, req *serializers.GetPagingExamReq) ([]*models.Exam, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	query := r.db
	order := "created_at"
	if req.Name != "" {
		query = query.
			Where("name LIKE ?", "%"+req.Name+"%").
			Where("user_id = ?", req.UserID)
	}
	if req.SubjectID != 0 {
		query = query.
			Where("subject_id = ?", req.SubjectID)
	}
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}
	var total int64
	if err := query.Model(&models.Exam{}).Count(&total).Error; err != nil {
		return nil, nil, errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var exams []*models.Exam
	if err := query.
		Preload("Questions").
		Limit(int(pagination.Limit)).
		Offset(int(pagination.Skip)).
		Order(order).
		Find(&exams).Error; err != nil {
		return nil, nil, nil
	}

	return exams, pagination, nil
}

func (r *ExamRepo) GetAll(ctx context.Context, userID uint) ([]*models.Exam, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var exams []*models.Exam
	if err := r.db.Where("user_id = ?", userID).Find(&exams).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return exams, nil
}

func (r *ExamRepo) GetByID(ctx context.Context, id uint, userID uint) (*models.Exam, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var exam models.Exam
	if err := r.db.Where("id = ?", id).Where("user_id = ?", userID).First(&exam).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &exam, nil
}

func (r *ExamRepo) Create(ctx context.Context, exam *models.Exam) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Where("name = ?", exam.Name).Where("user_id = ?", exam.UserID).First(&exam).Error; err == nil {
		return errors.ErrorExistName.New()
	}

	if err := r.db.Create(&exam).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *ExamRepo) Update(ctx context.Context, exam *models.Exam) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Where("name = ?", exam.Name).Where("user_id = ?", exam.UserID).First(&exam).Error; err == nil {
		return errors.ErrorExistName.New()
	}

	if err := r.db.Save(&exam).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *ExamRepo) Delete(ctx context.Context, exam *models.Exam) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	rowsAffected := r.db.Delete(&exam).RowsAffected

	if rowsAffected == 0 {
		return errors.ErrorNotFound.New()
	}

	return nil
}
