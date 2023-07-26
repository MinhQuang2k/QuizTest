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
	"quiztest/pkg/utils"
)

type IQuestionRepository interface {
	Create(ctx context.Context, question *models.Question) error
	Clones(ctx context.Context, userID uint, questionClonesID uint) error
	Update(ctx context.Context, question *models.Question) error
	GetPaging(ctx context.Context, req *serializers.GetPagingQuestionReq) ([]*models.Question, *paging.Pagination, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Question, error)
	Delete(ctx context.Context, question *models.Question) error
}

type QuestionRepo struct {
	db *gorm.DB
}

func NewQuestionRepository() *QuestionRepo {
	return &QuestionRepo{db: dbs.Database}
}

func (r *QuestionRepo) GetPaging(ctx context.Context, req *serializers.GetPagingQuestionReq) ([]*models.Question, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	query := r.db
	order := "created_at"
	if req.Name != "" {
		query = query.
			Where("name LIKE ?", "%"+req.Name+"%").
			Where("user_id = ?", req.UserID)
	}
	if req.GroupQuestionID != 0 {
		query = query.Where("group_question_id = ?", req.GroupQuestionID)
	}
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}
	var total int64
	if err := query.Model(&models.Question{}).Count(&total).Error; err != nil {
		return nil, nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var questions []*models.Question
	if err := query.
		Limit(int(pagination.Limit)).
		Offset(int(pagination.Skip)).
		Order(order).
		Find(&questions).Error; err != nil {
		return nil, nil, nil
	}

	return questions, pagination, nil
}

func (r *QuestionRepo) GetByID(ctx context.Context, id uint, userID uint) (*models.Question, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var question models.Question
	if err := r.db.Where("id = ?", id).Where("user_id = ?", userID).First(&question).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &question, nil
}

func (r *QuestionRepo) Create(ctx context.Context, question *models.Question) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Create(&question).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *QuestionRepo) Clones(ctx context.Context, userID uint, questionClonesID uint) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var question models.Question
	if err := r.db.Where("id = ?", questionClonesID).First(&question).Error; err != nil {
		return errors.ErrorDatabaseGet.Newm(err.Error())
	}
	var questionClones serializers.QuestionClones

	utils.Copy(&questionClones, &question)
	questionClones.UserID = userID

	var questionNew models.Question
	utils.Copy(&questionNew, &questionClones)
	if err := r.db.Create(&questionNew).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *QuestionRepo) Update(ctx context.Context, question *models.Question) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.Where("name = ?", question.Name).Where("user_id = ?", question.UserID).First(&question).Error; err == nil {
		return errors.ErrorExistName.New()
	}

	if err := r.db.Save(&question).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *QuestionRepo) Delete(ctx context.Context, question *models.Question) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	rowsAffected := r.db.Delete(&question).RowsAffected

	if rowsAffected == 0 {
		return errors.ErrorNotFound.New()
	}

	return nil
}
