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
	CategoryID  uint   `json:"category_id" gorm:"type:int;not null;"`
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

type CompetitionWithCategory struct {
	ID           uint   `json:"id"`
	Title        string `json:"title" gorm:"type:varchar(255);not null;"`
	EnrollTime   string `json:"enrolltime" gorm:"type:varchar(255);not null;"`
	StartTime    string `json:"starttime" gorm:"type:varchar(255);not null;"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"categoryname"`
}

type CompetitionNoCategory struct {
	ID         uint   `json:"id"`
	Title      string `json:"title" gorm:"type:varchar(255);not null;"`
	EnrollTime string `json:"enrolltime" gorm:"type:varchar(255);not null;"`
	StartTime  string `json:"starttime" gorm:"type:varchar(255);not null;"`
}

type GetCompetitionResponse struct {
	Competitions []CompetitionNoCategory `json:"competitions"`
	CategoryID   uint                    `json:"category_id"`
	CategoryName string                  `json:"categoryname"`
}
