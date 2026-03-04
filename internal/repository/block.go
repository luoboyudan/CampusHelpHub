package repository

import (
	"campushelphub/model"
	"context"

	"gorm.io/gorm"
)

type BlockRepository interface {
	CreateBlock(ctx context.Context, block *model.Block) error
	GetAllBlock(ctx context.Context) ([]model.Block, error)
}

type MySQLBlockRepository struct {
	db *gorm.DB
}

func NewMySQLBlockRepository(db *gorm.DB) *MySQLBlockRepository {
	return &MySQLBlockRepository{
		db: db,
	}
}

func (r *MySQLBlockRepository) CreateBlock(ctx context.Context, block *model.Block) error {
	return r.db.WithContext(ctx).Create(block).Error
}

func (r *MySQLBlockRepository) GetAllBlock(ctx context.Context) ([]model.Block, error) {
	var blocks []model.Block
	if err := r.db.WithContext(ctx).Find(&blocks).Error; err != nil {
		return nil, err
	}
	return blocks, nil
}
