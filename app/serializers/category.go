package serializers

import (
	"quiztest/pkg/paging"
	"time"
)

type Category struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Subjects  []*Subject `json:"subjects"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type GetPagingCategoryReq struct {
	UserID    string `json:"user_id" validate:"required"`
	Name      string `json:"name,omitempty" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
}

type GetPagingCategoryRes struct {
	Categories []*Category        `json:"rows"`
	Pagination *paging.Pagination `json:"pagination"`
}

type CreateCategoryReq struct {
	UserID   string     `json:"user_id" validate:"required"`
	Name     string     `json:"name" validate:"required"`
	Subjects []*Subject `json:"subjects"`
}
type CreateCategoryRes struct {
	Name     string     `json:"name" validate:"required"`
	Subjects []*Subject `json:"subjects"`
}

type UpdateCategoryReq struct {
	UserID string `json:"user_id" validate:"required"`
	Name   string `json:"name,omitempty"`
}
