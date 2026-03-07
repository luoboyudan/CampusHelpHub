package repository

import (
	"campushelphub/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type ReminderRepository interface {
	CreateReminder(ctx context.Context, reminder *model.Reminder) error
	GetRemindersToday() ([]*model.Reminder, error)
}

type MySQLReminderRepository struct {
	db *gorm.DB
}

func NewMySQLReminderRepository(db *gorm.DB) ReminderRepository {
	return &MySQLReminderRepository{db: db}
}

func (r *MySQLReminderRepository) CreateReminder(ctx context.Context, reminder *model.Reminder) error {
	return r.db.WithContext(ctx).Create(reminder).Error
}

func (r *MySQLReminderRepository) GetRemindersToday() ([]*model.Reminder, error) {
	var reminders []*model.Reminder
	if err := r.db.
		Where("reminder_time = ?", time.Now().Format("2006-01-02")).
		Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}
