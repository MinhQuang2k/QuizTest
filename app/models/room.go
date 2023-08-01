package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
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
