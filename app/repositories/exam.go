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
	Create(ctx context.Context, exam *models.Exam, userID uint) error
	Update(ctx context.Context, exam *models.Exam) error
	UpdateExamQuestion(ctx context.Context, examQuestion *models.ExamQuestion) error
	GetPaging(ctx context.Context, req *serializers.GetPagingExamReq) ([]*models.Exam, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Exam, error)
	GetExamQuestionByID(ctx context.Context, examID uint, questionID uint) (*models.ExamQuestion, error)
	GetAll(ctx context.Context, userID uint) ([]*models.Exam, error)
	Delete(ctx context.Context, exam *models.Exam) error
	DeleteExamQuestion(ctx context.Context, examQuestion *models.ExamQuestion) error
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
	var total int64
	var exams []*models.Exam

	query := r.db.Preload("Rooms").
		Joins("JOIN subjects ON subjects.id = exams.subject_id").
		Joins("JOIN categories ON categories.id = subjects.category_id").
		Where("categories.user_id = ?", req.UserID)

	order := "exams.created_at"
	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.SubjectID != 0 {
		query = query.Where("subject_id = ?", req.SubjectID)
	}
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	pagination := paging.New(req.Page, req.Limit, total)

	if err := query.
		Limit(int(pagination.Limit)).
		Offset(int(pagination.Skip)).
		Order(order).
		Find(&exams).
		Count(&total).Error; err != nil {
		return nil, nil, nil
	}

	pagination.Total = total

	return exams, pagination, nil
}

func (r *ExamRepo) GetAll(ctx context.Context, userID uint) ([]*models.Exam, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var exams []*models.Exam
	if err := r.db.
		Joins("JOIN subjects ON subjects.id = exams.subject_id").
		Joins("JOIN categories ON categories.id = subjects.category_id").
		Where("categories.user_id = ?", userID).Find(&exams).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return exams, nil
}

func (r *ExamRepo) GetByID(ctx context.Context, id uint, userID uint) (*models.Exam, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var exam models.Exam
	if err := r.db.Preload("Rooms").
		Joins("JOIN subjects ON subjects.id = exams.subject_id").
		Joins("JOIN categories ON categories.id = subjects.category_id").
		Where("categories.user_id = ?", userID).
		Where("exams.id = ?", id).
		First(&exam).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &exam, nil
}

func (r *ExamRepo) Create(ctx context.Context, exam *models.Exam, userID uint) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Create(&exam).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *ExamRepo) Update(ctx context.Context, exam *models.Exam) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Save(&exam).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *ExamRepo) UpdateExamQuestion(ctx context.Context, examQuestion *models.ExamQuestion) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Save(&examQuestion).Error; err != nil {
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

func (r *ExamRepo) GetExamQuestionByID(ctx context.Context, examID uint, questionID uint) (*models.ExamQuestion, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var examQuestion models.ExamQuestion
	if err := r.db.Where("exam_id = ?", examID).Where("question_id = ?", questionID).First(&examQuestion).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &examQuestion, nil
}

func (r *ExamRepo) DeleteExamQuestion(ctx context.Context, examQuestion *models.ExamQuestion) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	rowsAffected := r.db.Delete(&examQuestion).RowsAffected

	if rowsAffected == 0 {
		return errors.ErrorNotFound.New()
	}

	return nil
}
