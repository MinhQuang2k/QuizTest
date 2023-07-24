package serializers

import (
	"quiztest/pkg/paging"
	"time"
)

type GroupQuestion struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetPagingGroupQuestionReq struct {
	UserID    string `json:"user_id" validate:"required"`
	Name      string `json:"name,omitempty" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
}

type GetPagingGroupQuestionRes struct {
	GroupQuestions []*GroupQuestion   `json:"rows"`
	Pagination     *paging.Pagination `json:"pagination"`
}

type CreateGroupQuestionReq struct {
	UserID string `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type UpdateGroupQuestionReq struct {
	UserID string `json:"user_id" validate:"required"`
	Name   string `json:"name,omitempty"`
}
