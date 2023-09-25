package models

import "gorm.io/gorm"

type Exam struct {
	gorm.Model
	Name            string   `json:"name"`
	TimeLimit       uint     `json:"time_limit"`
	Description     string   `json:"description"`
	ShuffleQuestion bool     `json:"shuffle_question" gorm:"default:false"`
	Type            int      `json:"type"`
	SubjectID       uint     `json:"subject_id"`
	Subject         *Subject `json:"subject"`
	ExamQuestions   []*ExamQuestion
	Rooms           []*Room
}

type ExamQuestion struct {
	gorm.Model
	ExamID     uint `json:"exam_id"`
	QuestionID uint `json:"question_id"`
	Question   *Question
}
