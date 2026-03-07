package repository

import (
	"campushelphub/model"
	"context"

	"gorm.io/gorm"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *model.Message) error
}

type MySQLMessageRepository struct {
	db *gorm.DB
}

func NewMySQLMessageRepository(db *gorm.DB) MessageRepository {
	return &MySQLMessageRepository{db: db}
}

func (r *MySQLMessageRepository) CreateMessage(ctx context.Context, message *model.Message) error {
	return r.db.WithContext(ctx).Create(message).Error
}
