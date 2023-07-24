package serializers

import (
	"time"
)

type Subject struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSubjectReq struct {
	CategoryID string `json:"category_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
}

type UpdateSubjectReq struct {
	UserID     string `json:"user_id" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
	Name       string `json:"name,omitempty"`
}

type MoveSubjectReq struct {
	UserID        string `json:"user_id" validate:"required"`
	CategoryID    string `json:"category_id" validate:"required"`
	NewCategoryID string `json:"new_category_id" validate:"required"`
}
