package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subject struct {
	Base
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
	Category   *Category
}

func (subject *Subject) BeforeCreate(tx *gorm.DB) error {
	subject.ID = uuid.New().String()
	return nil
}
