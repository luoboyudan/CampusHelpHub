package model

import (
	"gorm.io/gorm"
)

// 竞赛
type Competition struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(255);not null;"`
	Description string `json:"description" gorm:"type:varchar(255);not null;"`
	EnrollTime  string `json:"enrolltime" gorm:"type:varchar(255);not null;"`
	StartTime   string `json:"starttime" gorm:"type:varchar(255);not null;"`
	Website     string `json:"website" gorm:"type:varchar(255);not null;"`
	Experience  string `json:"experience" gorm:"type:varchar(255);not null;"`
	BlockID     uint   `json:"blockid" gorm:"type:int;not null;"`
}

func (Competition) TableName() string {
	return "competitions"
}

type CreateCompetitionRequest struct {
	Competition
}

type CreateCompetitionResponse struct {
	Result bool `json:"result"`
}

type CompetitionWithBlock struct {
	ID         uint   `json:"id"`
	Title      string `json:"title" gorm:"type:varchar(255);not null;"`
	EnrollTime string `json:"enrolltime" gorm:"type:varchar(255);not null;"`
	StartTime  string `json:"starttime" gorm:"type:varchar(255);not null;"`
	BlockID    uint   `json:"blockid"`
	BlockName  string `json:"blockname"`
}

type CompetitionNoBlock struct {
	ID         uint   `json:"id"`
	Title      string `json:"title" gorm:"type:varchar(255);not null;"`
	EnrollTime string `json:"enrolltime" gorm:"type:varchar(255);not null;"`
	StartTime  string `json:"starttime" gorm:"type:varchar(255);not null;"`
}

type GetCompetitionResponse struct {
	Competitions []CompetitionNoBlock `json:"competitions"`
	BlockID      uint                 `json:"blockid"`
	BlockName    string               `json:"blockname"`
}
