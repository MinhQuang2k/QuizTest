package models

import (
	"gorm.io/gorm"
)

type Exam struct {
	gorm.Model
	Name            string `json:"name"`
	TimeLimit       uint   `json:"time_limit"`
	Description     string `json:"description"`
	ShuffleQuestion bool   `json:"shuffle_question" gorm:"default:false"`
	SubjectID       uint   `json:"subject_id"`
	Subject         *Subject
	Questions       []uint `gorm:"type:json"`
}
