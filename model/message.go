package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Title               string          `json:"title" gorm:"type:varchar(255);not null;"`
	Abstract            string          `json:"abstract" gorm:"type:text;default:''"`
	Content             string          `json:"content" gorm:"type:text;not null;"`
	Category            MessageCategory `json:"category" gorm:"type:varchar(50);not null;index:category_index;"`
	AuthorID            uint            `json:"author_id" gorm:"type:uint;not null;"`
	CoverImage          string          `json:"cover_image" gorm:"type:varchar(500);not null;"`
	IsTop               bool            `json:"is_top" gorm:"type:boolean;default:false"`
	ViewCount           int             `json:"view_count" gorm:"type:int;default:0"`
	Remind              bool            `json:"remind" gorm:"type:boolean;default:false"`
	RemindTime          time.Time       `json:"remind_time" gorm:"type:timestamp;default:current_timestamp"`
	RemindCompetitionID uint            `json:"remind_competition_id" gorm:"type:uint;default:0"`
	Status              MessageStatus   `json:"status" gorm:"type:varchar(50);not null;"`
	PublishedTime       time.Time       `json:"published_time" gorm:"type:timestamp;default:current_timestamp;index:published_time_index;"`
}

// Category使用枚举值
type MessageCategory string

const (
	CategoryCompetitionNotice MessageCategory = "competition_notice" // 竞赛通知
	CategoryTeamRecruitment   MessageCategory = "team_recruitment"   // 团队招募
	CategoryResultPublic      MessageCategory = "result_public"      // 结果公示
	CategoryAwardNews         MessageCategory = "award_news"         // 获奖喜报
	CategoryAcademicLecture   MessageCategory = "academic_lecture"   // 学术讲座
	CategoryActivity          MessageCategory = "activity"           // 文体活动/志愿服务
	CategoryGeneral           MessageCategory = "general"            // 学院综合通知
)

// Status使用枚举值
type MessageStatus string

const (
	StatusPublished MessageStatus = "published" // 已发布
	StatusDraft     MessageStatus = "draft"     // 草稿
)

type CreateMessageRequest struct {
	Title               string          `json:"title" binding:"required"`
	Content             string          `json:"content" binding:"required"`
	Category            MessageCategory `json:"category" binding:"required"`
	AuthorID            uint            `json:"author_id" binding:"required"`
	CoverImage          string          `json:"cover_image"`
	IsTop               bool            `json:"is_top"`
	Remind              bool            `json:"remind"`
	RemindUserID        uint            `json:"remind_user_id"`
	RemindCompetitionID uint            `json:"remind_competition_id"`
	Status              MessageStatus   `json:"status" binding:"required"`
}

type CreateMessageResponse struct {
	Result bool `json:"result"`
}
