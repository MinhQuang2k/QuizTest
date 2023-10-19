package repositories

import (
	"context"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/config"
	"quiztest/pkg/errors"
	"quiztest/pkg/paging"
	"quiztest/pkg/utils"
)

type QuestionRepo struct {
	db interfaces.IDatabase
}

func NewQuestionRepository(db interfaces.IDatabase) interfaces.IQuestionRepository {
	return &QuestionRepo{db: db}
}

func (r *QuestionRepo) GetPaging(ctx context.Context, req *serializers.GetPagingQuestionReq) ([]*models.Question, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()
	var total int64
	var questions []*models.Question
	query := r.db.GetInstance().
		Preload("GroupQuestion").
		Joins("JOIN group_questions ON group_questions.id = questions.group_question_id").
		Where("group_questions.user_id = ?", req.UserID)

	order := "questions.created_at DESC"
	if req.Content != "" {
		query = query.
			Where("questions.content LIKE ?", "%"+req.Content+"%")
	}
	if req.GroupQuestionID != 0 {
		query = query.Where("questions.group_question_id = ?", req.GroupQuestionID)
	}

	if req.SortBy == SortBy["RECENT"] {
		order = "questions.created_at DESC"
	}

	if req.SortBy == SortBy["ALPHABET"] {
		order = "questions.content"
	}

	pagination := paging.New(req.Page, req.Limit, total)
	if err := query.
		Limit(int(pagination.Limit)).
		Offset(int(pagination.Skip)).
		Order(order).
		Find(&questions).
		Count(&total).Error; err != nil {
		return nil, nil, nil
	}
	pagination.Total = total

	return questions, pagination, nil
}

func (r *QuestionRepo) GetByID(ctx context.Context, id uint, userID uint) (*models.Question, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var question models.Question
	if err := r.db.GetInstance().
		Joins("JOIN group_questions ON group_questions.id = questions.group_question_id").
		Where("questions.id = ?", id).
		Where("questions.deleted_at IS NULL").
		Where("group_questions.user_id = ?", userID).
		First(&question).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return &question, nil
}

func (r *QuestionRepo) GetByExamID(ctx context.Context, examID uint) ([]*models.Question, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var questions []*models.Question
	if err := r.db.GetInstance().
		Joins("JOIN exam_questions ON exam_questions.question_id = questions.id").
		Where("exam_questions.exam_id = ?", examID).
		Where("exam_questions.deleted_at IS NULL").
		Order("questions.created_at DESC").
		Find(&questions).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}

	return questions, nil
}

func (r *QuestionRepo) Create(ctx context.Context, question *models.Question) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.GetInstance().Create(&question).Error; err != nil {
		return errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return nil
}

func (r *QuestionRepo) Clones(ctx context.Context, userID uint, questionClonesID uint) (*models.Question, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var question models.Question
	if err := r.db.GetInstance().Where("id = ?", questionClonesID).First(&question).Error; err != nil {
		return nil, errors.ErrorDatabaseGet.Newm(err.Error())
	}
	var questionClones serializers.QuestionClones

	utils.Copy(&questionClones, &question)
	questionClones.UserID = userID

	var questionNew models.Question
	utils.Copy(&questionNew, &questionClones)
	if err := r.db.GetInstance().Create(&questionNew).Error; err != nil {
		return nil, errors.ErrorDatabaseCreate.Newm(err.Error())
	}

	return &questionNew, nil
}

func (r *QuestionRepo) Update(ctx context.Context, question *models.Question, userID uint) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.GetInstance().Save(&question).Error; err != nil {
		return errors.ErrorDatabaseUpdate.Newm(err.Error())
	}

	return nil
}

func (r *QuestionRepo) Delete(ctx context.Context, question *models.Question) error {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	if err := r.db.GetInstance().Where("question_id = ?", question.Model.ID).Delete(&models.ExamQuestion{}); err == nil {
		return errors.ErrorDatabaseDelete.New()
	}

	rowsAffected := r.db.GetInstance().Delete(&question).RowsAffected

	if rowsAffected == 0 {
		return errors.ErrorDatabaseDelete.New()
	}

	return nil
}

func (r *QuestionRepo) GetTotalScore(ctx context.Context, questions []uint) uint {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeout)
	defer cancel()

	var result struct {
		totalScore int
	}
	if len(questions) == 0 {
		return 0
	}
	if err := r.db.GetInstance().
		Table("questions").
		Select("SUM(score) as totalScore").
		Where("id IN (?)", questions).
		Find(&result.totalScore).Error; err != nil {
		return 0
	}

	return uint(result.totalScore)
}
