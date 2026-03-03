package model

import (
	"gorm.io/gorm"
)

// 用户
type User struct {
	gorm.Model
	ID        uint64 `json:"id" gorm:"primaryKey"` //雪花ID
	OpenID    string `json:"openid" gorm:"type:varchar(255);not null;unique"`
	StudentID string `json:"studentid" gorm:"type:varchar(255);not null;unique"`
	Username  string `json:"username" gorm:"type:varchar(255);not null"`
	Avatar    string `json:"avatar" gorm:"type:varchar(255);"`
	Bio       string `json:"bio" gorm:"type:varchar(255);"`
	School    string `json:"school" gorm:"type:varchar(255);"`
	Auth      bool   `json:"auth" gorm:"type:boolean;default:false;"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

type CreateUserRequest struct {
	Code     string `json:"code" binding:"required"`
	Username string `json:"username" binding:"required"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	School   string `json:"school"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type VerifyUserRequest struct {
	StudentID   string `json:"studentid" binding:"required"`
	RSAPassword string `json:"password" binding:"required"`
	UserID      uint64 `json:"userid" binding:"omitempty"`
}
