package model

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	ID        uint64 `json:"id" gorm:"primaryKey"` //雪花ID
	UserID    uint64 `json:"user_id" gorm:"type:integer;"`
	OrderID   uint64 `json:"order_id" gorm:"type:integer;"`
	Name      string `json:"name" gorm:"type:varchar(255);not null;"`
	District  string `json:"district" gorm:"type:varchar(255);not null;"`
	Address   string `json:"address" gorm:"type:varchar(255);not null;"`
	Latitude  string `json:"latitude" gorm:"type:varchar(255);not null;"`
	Longitude string `json:"longitude" gorm:"type:varchar(255);not null;"`
}

func (Location) TableName() string {
	return "locations"
}
