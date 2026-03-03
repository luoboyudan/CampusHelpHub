package model

import (
	"gorm.io/gorm"
)

// 竞赛
type Competition struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);not null;"`
	Description string `json:"description" gorm:"type:varchar(255);not null;"`
	EnrollTime  uint64 `json:"enrolltime" gorm:"type:bigint;not null;"`
	StartTime   uint64 `json:"starttime" gorm:"type:bigint;not null;"`
	Website     string `json:"website" gorm:"type:varchar(255);not null;"`
	Experience  string `json:"experience" gorm:"type:varchar(255);not null;"`
	BlockID     uint   `json:"blockid" gorm:"type:int;not null;"`
}

func (Competition) TableName() string {
	return "competitions"
}

type CreateCompetitionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	EnrollTime  uint64 `json:"enrolltime" binding:"omitempty"`
	StartTime   uint64 `json:"starttime" binding:"omitempty"`
	Website     string `json:"website" binding:"omitempty"`
	Experience  string `json:"experience" binding:"omitempty"`
}

type CreateCompetitionResponse struct {
}
