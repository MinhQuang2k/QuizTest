package serializers

import (
	"quiztest/pkg/paging"
	"time"
)

type Exam struct {
	ID              uint        `json:"id"`
	Name            string      `json:"name"`
	TimeLimit       uint        `json:"time_limit"`
	Description     string      `json:"description"`
	ShuffleQuestion bool        `json:"shuffle_question"`
	SubjectID       uint        `json:"subject_id"`
	UserID          uint        `json:"user_id"`
	Questions       []*Question `json:"questions"`
	Rooms           []*Room     `json:"rooms"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type GetPagingExamReq struct {
	UserID    uint   `json:"user_id" validate:"required"`
	SubjectID uint   `json:"subject_id,omitempty" form:"subject_id"`
	Name      string `json:"name,omitempty" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
}

type GetPagingExamRes struct {
	Exams      []*Exam            `json:"rows"`
	Pagination *paging.Pagination `json:"pagination"`
}

type CreateExamReq struct {
	UserID          uint   `json:"user_id" validate:"required"`
	Name            string `json:"name"`
	TimeLimit       uint   `json:"time_limit"`
	Description     string `json:"description"`
	ShuffleQuestion bool   `json:"shuffle_question"`
	SubjectID       uint   `json:"subject_id"`
}

type UpdateExamReq struct {
	UserID          uint   `json:"user_id" validate:"required"`
	Name            string `json:"name"`
	TimeLimit       uint   `json:"time_limit"`
	Description     string `json:"description"`
	ShuffleQuestion bool   `json:"shuffle_question"`
	SubjectID       uint   `json:"subject_id"`
}

type MoveExamReq struct {
	UserID         uint `json:"user_id" validate:"required"`
	ExamID         uint `json:"exam_id" validate:"required"`
	QuestionID     uint `json:"question_id" validate:"required"`
	QuestionMoveID uint `json:"question_move_id" validate:"required"`
}
