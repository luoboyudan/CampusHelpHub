package repository

import (
	"campushelphub/model"
	"context"

	"gorm.io/gorm"
)

type CompetitionRepository interface {
	CreateCompetition(ctx context.Context, competition *model.Competition) error
	GetCompetitionByBlockID(ctx context.Context, blockID uint64) ([]model.Competition, error)
	GetCompetitions(ctx context.Context) ([]model.CompetitionWithBlock, error)
}

type MySQLCompetitionRepository struct {
	db *gorm.DB
}

func NewMySQLCompetitionRepository(db *gorm.DB) CompetitionRepository {
	return &MySQLCompetitionRepository{
		db: db,
	}
}

func (r *MySQLCompetitionRepository) CreateCompetition(ctx context.Context, competition *model.Competition) error {
	return r.db.WithContext(ctx).Create(competition).Error
}

func (r *MySQLCompetitionRepository) GetCompetitionByBlockID(ctx context.Context, blockID uint64) ([]model.Competition, error) {
	var competitions []model.Competition
	if err := r.db.WithContext(ctx).Where("block_id = ?", blockID).Find(&competitions).Error; err != nil {
		return nil, err
	}
	return competitions, nil
}

func (r *MySQLCompetitionRepository) GetCompetitions(ctx context.Context) ([]model.CompetitionWithBlock, error) {
	var comps []model.CompetitionWithBlock
	if err := r.db.WithContext(ctx).Select("competitions.*, board.name as block_name").Joins("JOIN board ON competitions.block_id = board.id").Find(&comps).Error; err != nil {
		return nil, err
	}
	return comps, nil
}
