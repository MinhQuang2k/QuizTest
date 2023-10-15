package serializers

import (
	"quiztest/pkg/paging"
	"time"

	"gorm.io/datatypes"
)

type Room struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name"`
	AttemptLimit uint           `json:"attempt_limit"`
	TypeCode     uint           `json:"type_code"`
	PassMark     uint           `json:"pass_mark"`
	AccessCodes  datatypes.JSON `json:"access_codes" gorm:"type:json"`
	Requires     datatypes.JSON `json:"requires" gorm:"type:json"`
	IsActive     bool           `json:"is_active" gorm:"default:false"`
	LinkRoomId   string         `json:"link_room_id"`
	ScoreShown   datatypes.JSON `json:"score_shown" gorm:"type:json"`
	ResultShown  datatypes.JSON `json:"result_shown" gorm:"type:json"`
	EndAt        string         `json:"end_at"`
	StartAt      string         `json:"start_at"`
	Note         string         `json:"note"`
	ExamID       uint           `json:"exam_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type GetPagingRoomReq struct {
	UserID    uint   `json:"user_id" validate:"required"`
	Name      string `json:"name,omitempty" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
}

type GetCodeRoomReq struct {
	Limit int `json:"limit" form:"limit"`
}

type GetPagingRoomRes struct {
	Rooms      []*Room            `json:"rows"`
	Pagination *paging.Pagination `json:"pagination"`
}

type CreateRoomReq struct {
	UserID       uint           `json:"user_id" validate:"required"`
	Name         string         `json:"name"`
	AttemptLimit uint           `json:"attempt_limit"`
	TypeCode     uint           `json:"type_code"`
	PassMark     uint           `json:"pass_mark"`
	AccessCodes  datatypes.JSON `json:"access_codes" gorm:"type:json"`
	Requires     datatypes.JSON `json:"requires" gorm:"type:json"`
	IsActive     bool           `json:"is_active" gorm:"default:false"`
	LinkRoomId   string         `json:"link_room_id"`
	ScoreShown   datatypes.JSON `json:"score_shown" gorm:"type:json"`
	ResultShown  datatypes.JSON `json:"result_shown" gorm:"type:json"`
	EndAt        string         `json:"end_at"`
	StartAt      string         `json:"start_at"`
	Note         string         `json:"note"`
	ExamID       uint           `json:"exam_id"`
}

type UpdateRoomReq struct {
	UserID       uint           `json:"user_id" validate:"required"`
	Name         string         `json:"name"`
	AttemptLimit uint           `json:"attempt_limit"`
	TypeCode     uint           `json:"type_code"`
	PassMark     uint           `json:"pass_mark"`
	AccessCodes  datatypes.JSON `json:"access_codes" gorm:"type:json"`
	Requires     datatypes.JSON `json:"requires" gorm:"type:json"`
	IsActive     bool           `json:"is_active" gorm:"default:false"`
	LinkRoomId   string         `json:"link_room_id"`
	ScoreShown   datatypes.JSON `json:"score_shown" gorm:"type:json"`
	ResultShown  datatypes.JSON `json:"result_shown" gorm:"type:json"`
	EndAt        string         `json:"end_at"`
	StartAt      string         `json:"start_at"`
	Note         string         `json:"note"`
	ExamID       uint           `json:"exam_id"`
}
