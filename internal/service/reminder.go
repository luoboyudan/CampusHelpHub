package service

import (
	"campushelphub/internal/repository"
	"campushelphub/model"
	"context"
)

type ReminderService struct {
	ReminderRepository repository.ReminderRepository
	cron               *TaskService
}

func NewReminderService(reminderRepository repository.ReminderRepository, cron *TaskService) *ReminderService {
	return &ReminderService{
		ReminderRepository: reminderRepository,
		cron:               cron,
	}
}

func (s *ReminderService) CreateReminder(ctx context.Context, reminder *model.Reminder) error {
	if err := s.ReminderRepository.CreateReminder(ctx, reminder); err != nil {
		return err
	}
	return nil
}
