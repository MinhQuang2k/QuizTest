package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name   string `json:"name"`
	UserID string `json:"user_id"`
	User   *User
}

func (category *Category) BeforeCreate(tx *gorm.DB) error {
	category.ID = uuid.New().String()
	return nil
}
