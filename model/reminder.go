package model

import (
	"campushelphub/internal/pkg/wechat"
	"time"

	"gorm.io/gorm"
)

type Reminder struct {
	gorm.Model
	MessageID        uint                                 `json:"message_id" gorm:"type:uint;not null;"`
	UserOpenID       string                               `json:"user_open_id" gorm:"type:varchar(255);not null;"`
	Status           ReminderStatus                       `json:"status" gorm:"type:int;default:0"`
	WechatTemplateID string                               `json:"wechat_template_id" gorm:"type:varchar(255);default:''"`
	WechatPagePath   string                               `json:"wechat_page_path" gorm:"type:varchar(255);default:''"`
	Data             map[string]wechat.WechatReminderData `json:"data" gorm:"type:json;default:'{}'"`
	FailReason       string                               `json:"fail_reason" gorm:"type:text;default:''"`
	RemindTime       time.Time                            `json:"remind_time" gorm:"type:timestamp;default:current_timestamp"`
	SentAt           *time.Time                           `json:"sent_at" gorm:"type:timestamp;default:current_timestamp"`
}

// ReminderStatus 提醒状态枚举
type ReminderStatus int

const (
	ReminderStatusPending ReminderStatus = iota // 待发送
	ReminderStatusSent                          // 已发送
	ReminderStatusFailed                        // 发送失败
	ReminderStatusSkipped                       // 已跳过
)
