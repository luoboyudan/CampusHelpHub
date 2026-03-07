package model

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

// 竞赛
type Competition struct {
	gorm.Model
	Title       string              `json:"title" gorm:"type:varchar(255);not null;"`
	Description string              `json:"description" gorm:"type:text;not null;"`
	Editions    CompetitionEditions `json:"editions" gorm:"type:json;"`
	Website     string              `json:"website" gorm:"type:varchar(255);not null;"`
	Experience  string              `json:"experience" gorm:"type:varchar(255);not null;"`
	CategoryID  uint                `json:"category_id" gorm:"type:int;not null;foreignKey:CategoryID;references:ID;"`
}
type CompetitionEditions []CompetitionEdition

func (e CompetitionEditions) Value() (driver.Value, error) {
	if len(e) == 0 {
		return json.Marshal([]CompetitionEdition{})
	}
	return json.Marshal(e)
}
func (e *CompetitionEditions) Scan(value interface{}) error {
	if value == nil {
		*e = CompetitionEditions{}
		return nil
	}
	return json.Unmarshal([]byte(value.([]uint8)), &e)
}

type CompetitionEdition struct {
	Year      int                `json:"year" binding:"required"`
	Theme     []CompetitionTheme `json:"theme" binding:"omitempty"`
	Schedules []Schedules        `json:"schedules" binding:"required,dive"`
}

type CompetitionTheme struct {
	Name        string `json:"name" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
}
type Schedules struct {
	PhaseName string `json:"ph_name" binding:"required,min=1,max=255"`
	StartTime string `json:"start_time" binding:"required,datetime=2006-01-02 15:04:05"`
	Location  string `json:"location"`
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
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"categoryname"`
}

type CompetitionNoCategory struct {
	ID    uint   `json:"id"`
	Title string `json:"title" gorm:"type:varchar(255);not null;"`
}

type GetCompetitionResponse struct {
	Competitions []CompetitionNoCategory `json:"competitions"`
	CategoryID   uint                    `json:"category_id"`
	CategoryName string                  `json:"categoryname"`
}

type GetCompetitionRequest struct {
	ID uint `json:"id" binding:"required"`
}
