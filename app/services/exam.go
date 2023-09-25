package services

import (
	"context"

	"quiztest/pkg/errors"
	"quiztest/pkg/logger"

	"quiztest/app/interfaces"
	"quiztest/app/models"
	"quiztest/app/serializers"
	"quiztest/pkg/paging"
	"quiztest/pkg/utils"
)

type ExamService struct {
	repo         interfaces.IExamRepository
	repoQuestion interfaces.IQuestionRepository
}

func NewExamService(repo interfaces.IExamRepository, repoQuestion interfaces.IQuestionRepository) interfaces.IExamService {
	return &ExamService{repo: repo, repoQuestion: repoQuestion}
}

func (p *ExamService) GetByID(ctx context.Context, id uint, userID uint) (*serializers.Exam, []*models.Question, error) {
	exam, err := p.repo.GetByID(ctx, id, userID)

	if err != nil {
		return nil, nil, err
	}
	questions, err := p.repoQuestion.GetByExamID(ctx, exam.ID)

	var examsFormat serializers.Exam
	utils.Copy(&examsFormat, &exam)
	listQuestion := []uint{}
	for _, item := range questions {
		listQuestion = append(listQuestion, item.ID)
	}
	examsFormat.TotalQuestions = uint(len(listQuestion))
	examsFormat.TotalScore = p.repoQuestion.GetTotalScore(ctx, listQuestion)

	if err != nil {
		return nil, nil, err
	}

	return &examsFormat, questions, nil
}

func (p *ExamService) GetPaging(ctx context.Context, req *serializers.GetPagingExamReq) ([]*serializers.Exam, *paging.Pagination, error) {
	exams, pagination, err := p.repo.GetPaging(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	type arrayID struct {
		ID uint
	}
	var examsFormat []*serializers.Exam
	for _, ex := range exams {
		exFormat := &serializers.Exam{}
		utils.Copy(exFormat, ex)
		questions := []uint{}
		for _, item := range ex.ExamQuestions {
			questions = append(questions, item.ID)
		}
		exFormat.TotalQuestions = uint(len(questions))
		exFormat.TotalScore = p.repoQuestion.GetTotalScore(ctx, questions)
		examsFormat = append(examsFormat, exFormat)

	}
	return examsFormat, pagination, nil
}
func (p *ExamService) GetAll(ctx context.Context, userID uint) ([]*models.Exam, error) {
	exams, err := p.repo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	return exams, nil
}

func (p *ExamService) Create(ctx context.Context, req *serializers.CreateExamReq) error {
	var exam models.Exam
	utils.Copy(&exam, req)

	err := p.repo.Create(ctx, &exam, req.UserID)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return err
	}

	return nil
}

func (p *ExamService) Update(ctx context.Context, id uint, req *serializers.UpdateExamReq) error {
	exam, err := p.repo.GetByID(ctx, id, req.UserID)
	if err != nil {
		logger.Errorf("Update.GetByID fail, id: %s, error: %s", id, err)
		return err
	}

	utils.Copy(exam, req)
	err = p.repo.Update(ctx, exam)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (p *ExamService) Delete(ctx context.Context, id uint, userID uint) error {
	exam, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("Delete.GetByID fail, id: %s, error: %s", id, err)
		return err
	}

	err = p.repo.Delete(ctx, exam)
	if err != nil {
		logger.Errorf("Delete fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (p *ExamService) AddQuestion(ctx context.Context, id, question_id, userID uint) error {
	question, err := p.repoQuestion.GetByID(ctx, question_id, userID)
	if err != nil {
		logger.Errorf("AddQuestion.GetQuestion fail, id: %s, error: %s", question_id, err)
		return err
	}

	exam, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("AddQuestion.GetExam fail, id: %s, error: %s", id, err)
		return err
	}
	exitExamQuestion, _ := p.repo.GetExamQuestionByID(ctx, exam.ID, question.ID)

	if exitExamQuestion != nil {
		logger.Errorf("AddQuestion.Exam fail, id: %s, error: %s", id, err)
		return errors.ErrorExistName.New()
	}

	var examQuestion models.ExamQuestion
	examQuestion.QuestionID = question.ID
	exam.ExamQuestions = append(exam.ExamQuestions, &examQuestion)

	err = p.repo.Update(ctx, exam)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (p *ExamService) DeleteQuestion(ctx context.Context, id, question_id, userID uint) error {
	question, err := p.repoQuestion.GetByID(ctx, question_id, userID)
	if err != nil {
		logger.Errorf("AddQuestion.GetQuestion fail, id: %s, error: %s", question_id, err)
		return err
	}

	exam, err := p.repo.GetByID(ctx, id, userID)
	if err != nil {
		logger.Errorf("AddQuestion.GetExam fail, id: %s, error: %s", id, err)
		return err
	}
	examQuestion, err := p.repo.GetExamQuestionByID(ctx, exam.ID, question.ID)

	if err != nil {
		logger.Errorf("AddQuestion.Exam fail, id: %s, error: %s", id, err)
		return errors.ErrorExistName.New()
	}

	err = p.repo.DeleteExamQuestion(ctx, examQuestion)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (p *ExamService) MoveQuestion(ctx context.Context, req *serializers.MoveExamReq) error {
	examQuestion, err := p.repo.GetExamQuestionByID(ctx, req.ExamID, req.QuestionID)
	if err != nil {
		logger.Errorf("AddQuestion.Exam fail, id: %s, error: %s", req.QuestionID, err)
		return err
	}

	examQuestionMove, err := p.repo.GetExamQuestionByID(ctx, req.ExamID, req.QuestionMoveID)
	if err != nil {
		logger.Errorf("AddQuestion.Exam fail, id: %s, error: %s", req.QuestionMoveID, err)
		return err
	}

	examQuestion.QuestionID = req.QuestionMoveID
	examQuestionMove.QuestionID = req.QuestionID

	err = p.repo.UpdateExamQuestion(ctx, examQuestion)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", req.QuestionMoveID, err)
		return err
	}

	err = p.repo.UpdateExamQuestion(ctx, examQuestionMove)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", req.QuestionID, err)
		return err
	}

	return nil
}
