package service

import (
	"campushelphub/internal/config"
	"campushelphub/internal/errors"
	"campushelphub/internal/pkg/wechat"
	"campushelphub/internal/repository"

	"github.com/robfig/cron/v3"
)

type TaskService struct {
	cron               *cron.Cron
	reminderRepository repository.ReminderRepository
	wechatService      *wechat.WechatService
	config             *config.Config
	errs               *errors.Error
}

func NewTaskService(reminderRepository repository.ReminderRepository, wechatService *wechat.WechatService, config *config.Config, errs *errors.Error) *TaskService {
	taskService := TaskService{
		cron:               cron.New(),
		reminderRepository: reminderRepository,
		wechatService:      wechatService,
		config:             config,
		errs:               errs,
	}
	taskService.cron.Start()
	taskService.addReminderCheckTask()
	return &taskService
}

func (s *TaskService) checkReminder() error {
	reminders, err := s.reminderRepository.GetRemindersToday()
	if err != nil {
		return err
	}
	for _, reminder := range reminders {
		s.wechatService.SendReminder(&wechat.WechatReminder{
			Touser:     reminder.UserOpenID,
			TemplateID: reminder.WechatTemplateID,
			Page:       reminder.WechatPagePath,
			Data:       reminder.Data,
		})
	}
	return nil
}

// 每日检查一次提醒时间是否到，发送提醒
func (s *TaskService) addReminderCheckTask() error {
	s.cron.AddFunc(s.config.Task.ReminderTime, func() {
		s.checkReminder()
	})
	return nil
}
