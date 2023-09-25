package serializers

import (
	"quiztest/pkg/paging"
	"time"

	"gorm.io/datatypes"
)

type Question struct {
	ID                     uint           `json:"id"`
	Content                string         `json:"content"`
	Type                   int            `json:"type"`
	NoteAnswer             string         `json:"note_answer"`
	Answer                 datatypes.JSON `json:"answer" gorm:"type:json"`
	Score                  int            `json:"score" gorm:"default:0"`
	CorrectAnswer          datatypes.JSON `json:"correct_answer" gorm:"type:json"`
	MatchingCorrect        datatypes.JSON `json:"matching_correct" gorm:"type:json"`
	MatchingAnswer         datatypes.JSON `json:"matching_answer" gorm:"type:json"`
	FillBlankCorrectAnswer datatypes.JSON `json:"fill_blank_correct_answer" gorm:"type:json"`
	GroupQuestionID        uint           `json:"group_question_id"`
	GroupQuestion          *GroupQuestion
	UserID                 uint      `json:"user_id"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type QuestionClones struct {
	Content                string         `json:"content"`
	Type                   int            `json:"type"`
	NoteAnswer             string         `json:"note_answer"`
	Answer                 datatypes.JSON `json:"answer" gorm:"type:json"`
	Score                  int            `json:"score" gorm:"default:0"`
	CorrectAnswer          datatypes.JSON `json:"correct_answer" gorm:"type:json"`
	MatchingCorrect        datatypes.JSON `json:"matching_correct" gorm:"type:json"`
	MatchingAnswer         datatypes.JSON `json:"matching_answer" gorm:"type:json"`
	FillBlankCorrectAnswer datatypes.JSON `json:"fill_blank_correct_answer" gorm:"type:json"`
	GroupQuestionID        uint           `json:"group_question_id"`
	UserID                 uint           `json:"user_id"`
}

type GetPagingQuestionReq struct {
	UserID          uint   `json:"user_id" validate:"required"`
	GroupQuestionID uint   `json:"group_question_id" form:"group_question_id"`
	Content         string `json:"content,omitempty" form:"content"`
	Page            int64  `json:"-" form:"page"`
	Limit           int64  `json:"-" form:"limit"`
	SortBy          string `json:"-" form:"sort_by"`
}

type GetPagingQuestionRes struct {
	Questions  []*Question        `json:"rows"`
	Pagination *paging.Pagination `json:"pagination"`
}

type CreateQuestionReq struct {
	Content                string         `json:"content"`
	Type                   int            `json:"type"`
	NoteAnswer             string         `json:"note_answer"`
	Answer                 datatypes.JSON `json:"answer" gorm:"type:json"`
	Score                  int            `json:"score" gorm:"default:0"`
	CorrectAnswer          datatypes.JSON `json:"correct_answer" gorm:"type:json"`
	MatchingCorrect        datatypes.JSON `json:"matching_correct" gorm:"type:json"`
	MatchingAnswer         datatypes.JSON `json:"matching_answer" gorm:"type:json"`
	FillBlankCorrectAnswer datatypes.JSON `json:"fill_blank_correct_answer" gorm:"type:json"`
	GroupQuestionID        uint           `json:"group_question_id"`
	UserID                 uint           `json:"user_id"`
}

type UpdateQuestionReq struct {
	Content                string         `json:"content"`
	Type                   int            `json:"type"`
	NoteAnswer             string         `json:"note_answer"`
	Answer                 datatypes.JSON `json:"answer" gorm:"type:json"`
	Score                  int            `json:"score" gorm:"default:0"`
	CorrectAnswer          datatypes.JSON `json:"correct_answer" gorm:"type:json"`
	MatchingCorrect        datatypes.JSON `json:"matching_correct" gorm:"type:json"`
	MatchingAnswer         datatypes.JSON `json:"matching_answer" gorm:"type:json"`
	FillBlankCorrectAnswer datatypes.JSON `json:"fill_blank_correct_answer" gorm:"type:json"`
	GroupQuestionID        uint           `json:"group_question_id"`
	UserID                 uint           `json:"user_id"`
}
