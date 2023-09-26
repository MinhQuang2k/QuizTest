package serializers

import (
	"quiztest/pkg/paging"
	"time"
)

type Room struct {
	ID                    uint      `json:"id"`
	Name                  string    `json:"name"`
	PassMark              string    `json:"pass_mark"`
	IsRequireCode         bool      `json:"is_require_code"`
	IsRequireEmail        bool      `json:"is_require_email"`
	IsRequireFullName     bool      `json:"is_require_full_name"`
	IsRequirePhone        bool      `json:"is_require_phone"`
	IsRequireGroup        bool      `json:"is_require_group"`
	IsRequireIdentifyCode bool      `json:"is_require_identify_code"`
	CodeRoom              string    `json:"code_room"`
	LinkRoomExam          string    `json:"link_room_exam"`
	Status                bool      `json:"status"`
	ExamID                uint      `json:"exam_id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
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
	UserID                uint   `json:"user_id" validate:"required"`
	Name                  string `json:"name"`
	PassMark              string `json:"pass_mark"`
	IsRequireCode         bool   `json:"is_require_code"`
	IsRequireEmail        bool   `json:"is_require_email"`
	IsRequireFullName     bool   `json:"is_require_full_name"`
	IsRequirePhone        bool   `json:"is_require_phone"`
	IsRequireGroup        bool   `json:"is_require_group"`
	IsRequireIdentifyCode bool   `json:"is_require_identify_code"`
	CodeRoom              string `json:"code_room"`
	LinkRoomExam          string `json:"link_room_exam"`
	Status                bool   `json:"status"`
	ExamID                uint   `json:"exam_id"`
}

type UpdateRoomReq struct {
	UserID                uint   `json:"user_id" validate:"required"`
	Name                  string `json:"name"`
	PassMark              string `json:"pass_mark"`
	IsRequireCode         bool   `json:"is_require_code"`
	IsRequireEmail        bool   `json:"is_require_email"`
	IsRequireFullName     bool   `json:"is_require_full_name"`
	IsRequirePhone        bool   `json:"is_require_phone"`
	IsRequireGroup        bool   `json:"is_require_group"`
	IsRequireIdentifyCode bool   `json:"is_require_identify_code"`
	CodeRoom              string `json:"code_room"`
	LinkRoomExam          string `json:"link_room_exam"`
	Status                bool   `json:"status"`
	ExamID                uint   `json:"exam_id"`
}

// {
//     "name": "vfdvf",
//     "note": "vffvf",
//     "exam_id": 10342,
//     "password_type": "no_pass",
//     "password": null,
//     "attempt_limit": 0,
//     "name": "dcdcd",
//     "start_at": null,
//     "end_at": null,
//     "access_codes": [],
//     "is_require_phone": 0,
//     "is_require_email": 0,
//     "is_require_fullname": 1,
//     "is_require_identify_code": 0,
//     "is_require_group": 0,
//     "is_require_position": 0,
//     "status": true,
//     "link_room_exam": "https://e.testcenter.vn/t/cEJ3VH0OJFAOM1RcRyYED0F4SHIX",
//     "pass_mark": 80,
//     "is_score_shown": true,
//     "is_detail_result_shown": true,
//     "is_percent_shown": true,
//     "is_passed_result_shown": false
// }

// {
//     "name": "vfdvf",
//     "note": "vffvf",
//     "exam_id": 10342,
//     "start_at": null,
//     "end_at": null,
//     "type_code": 0,
//     "access_codes": [],
//     "requires": [],
//     "is_active": true,
//     "link_room_id": "cEJ3VH0OJFAOM1RcRyYED0F4SHIX",
//     "pass_mark": 80,
//     "score_shown": [],
//     "result_shown": []
// }
