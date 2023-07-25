package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Name                   string `json:"name"`
	Type                   string `json:"type"`
	NoteAnswer             string `json:"note_answer"`
	Answer                 string `json:"answer"`
	Score                  string `json:"score"`
	CorrectAnswer          string `json:"correct_answer"`
	HasMulCorrectAnswer    string `json:"has_mul_correct_answer"`
	MatchingCorrect        string `json:"matching_correct"`
	MatchingAnswer         string `json:"matching_answer"`
	FillBlankCorrectAnswer string `json:"fill_blank_correct_answer"`
	GroupQuestionID        uint   `json:"group_question_id"`
	GroupQuestion          *GroupQuestion
	UserID                 uint `json:"user_id"`
	User                   *User
}