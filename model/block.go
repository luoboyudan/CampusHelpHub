package model

import (
	"gorm.io/gorm"
)

// 竞赛板块
type Block struct {
	gorm.Model
	Name         string        `json:"name" gorm:"type:varchar(255);not null;"`
	Description  string        `json:"description" gorm:"type:varchar(255);not null;"`
	Competitions []Competition `json:"competitions" gorm:"foreignKey:BlockID;references:ID;"`
}

func (Block) TableName() string {
	return "blocks"
}

type CreateBlockRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateBlockResponse struct {
	Result bool `json:"result"`
	ID     uint `json:"id"`
}
