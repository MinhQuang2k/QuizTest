package serializers

import (
	"quiztest/pkg/paging"
)

type GroupQuestion struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetPagingGroupQuestionReq struct {
	UserID    uint   `json:"user_id" validate:"required"`
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
	UserID uint   `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type UpdateGroupQuestionReq struct {
	UserID uint   `json:"user_id" validate:"required"`
	Name   string `json:"name,omitempty"`
}
