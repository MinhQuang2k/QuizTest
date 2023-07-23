package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupQuestion struct {
	Base
	Name   string `json:"name"`
	UserID string `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID"`
}

func (groupQuestion *GroupQuestion) BeforeCreate(tx *gorm.DB) error {
	groupQuestion.ID = uuid.New().String()
	return nil
}
