package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID       uint64   `json:"id" gorm:"primaryKey"` //雪花ID
	UserID   uint64   `json:"user_id" gorm:"type:integer;"`
	Location Location `json:"location" gorm:"foreignKey:OrderID;references:ID;"`
}

func (Order) TableName() string {
	return "orders"
}
