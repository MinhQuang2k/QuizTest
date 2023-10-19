package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	FullName             string         `json:"full_name"`
	Email                string         `json:"email"`
	Phone                string         `json:"phone"`
	IdentifyCode         string         `json:"identify_code"`
	Group                string         `json:"group"`
	Position             string         `json:"position"`
	Answers              datatypes.JSON `json:"answer" gorm:"type:json"`
	TotalQuestion        uint           `json:"total_question"`
	TotalCorrectQuestion uint           `json:"total_correct_question"`
	TotalDoingQuestion   uint           `json:"total_doing_question"`
	MaxScore             uint           `json:"max_score"`
	Score                uint           `json:"score"`
	Status               string         `json:"status"`
	StartAt              string         `json:"start_at"`
	EndAt                string         `json:"end_at"`
	RoomId               uint           `json:"room_id"`
	Room                 *Room
}
