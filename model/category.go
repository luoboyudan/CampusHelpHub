package model

import (
	"gorm.io/gorm"
)

// 竞赛板块
type Category struct {
	gorm.Model
	Name         string        `json:"name" gorm:"type:varchar(255);not null;"`
	Description  string        `json:"description" gorm:"type:varchar(255);"`
	Competitions []Competition `json:"competitions" gorm:"foreignKey:BlockID;references:ID;"`
}

func (Category) TableName() string {
	return "categories"
}

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
}

type CreateCategoryResponse struct {
	Result bool `json:"result"`
	ID     uint `json:"id"`
}

type GetAllCategoryResponse struct {
	Categories []Category `json:"categories"`
}
