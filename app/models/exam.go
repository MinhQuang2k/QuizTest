package models

import (
	"gorm.io/gorm"
)

type Exam struct {
	gorm.Model
	Name            string `json:"name"`
	TimeLimit       string `json:"time_limit"`
	Description     string `json:"description"`
	ShuffleQuestion string `json:"shuffle_question"`
	SubjectID       uint   `json:"subject_id"`
	Subject         *Subject
	UserID          uint `json:"user_id"`
	User            *User
	Questions       []*Question `gorm:"many2many:exam_questions;"`
}
