package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
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
