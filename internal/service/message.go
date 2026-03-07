package service

import (
	"campushelphub/internal/pkg/wechat"
	"campushelphub/internal/repository"
	"campushelphub/model"
	"context"
)

type MessageService struct {
	MessageRepository  repository.MessageRepository
	ReminderRepository repository.ReminderRepository
}

func NewMessageService(messageRepository repository.MessageRepository, reminderRepository repository.ReminderRepository) *MessageService {
	return &MessageService{
		MessageRepository:  messageRepository,
		ReminderRepository: reminderRepository,
	}
}

func (s *MessageService) CreateMessage(ctx context.Context, message *model.Message) error {
	if message.Category == model.CategoryCompetitionNotice && message.Remind {
		reminder := &model.Reminder{
			MessageID:        message.ID,
			RemindTime:       message.RemindTime,
			WechatTemplateID: wechat.WechatTemplateIDCompetitionReminder,
			WechatPagePath:   wechat.WechatJumpURLMap[wechat.WechatTemplateIDCompetitionReminder],
		}
		if err := s.MessageRepository.CreateMessage(ctx, message); err != nil {
			return err
		}
		if err := s.ReminderRepository.CreateReminder(ctx, reminder); err != nil {
			return err
		}
	}
	return s.MessageRepository.CreateMessage(ctx, message)
}
