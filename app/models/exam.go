package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Exam struct {
	Base
	Name            string `json:"name"`
	TimeLimit       string `json:"time_limit"`
	Description     string `json:"description"`
	ShuffleQuestion string `json:"shuffle_question"`
	SubjectID       string `json:"subject_id"`
	Subject         *Subject
	UserID          string `json:"user_id"`
	User            *User
}

func (exam *Exam) BeforeCreate(tx *gorm.DB) error {
	exam.ID = uuid.New().String()
	return nil
}
