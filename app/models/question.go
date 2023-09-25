package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Content                string         `json:"content"`
	Type                   int            `json:"type"`
	NoteAnswer             string         `json:"note_answer"`
	Answer                 datatypes.JSON `json:"answer" gorm:"type:json"`
	Score                  int            `json:"score" gorm:"default:0"`
	CorrectAnswer          datatypes.JSON `json:"correct_answer" gorm:"type:json"`
	HasMulCorrectAnswer    bool           `json:"has_mul_correct_answer" gorm:"default:false"`
	MatchingCorrect        datatypes.JSON `json:"matching_correct" gorm:"type:json"`
	MatchingAnswer         datatypes.JSON `json:"matching_answer" gorm:"type:json"`
	FillBlankCorrectAnswer datatypes.JSON `json:"fill_blank_correct_answer" gorm:"type:json"`
	GroupQuestionID        uint           `json:"group_question_id"`
	GroupQuestion          *GroupQuestion
}
