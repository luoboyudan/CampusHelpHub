package repository

import (
	"campushelphub/model"
	"context"

	"gorm.io/gorm"
)

type CompetitionRepository interface {
	CreateCompetition(ctx context.Context, competition *model.Competition) error
	GetCompetitionByCategoryID(ctx context.Context, categoryID uint64) ([]model.Competition, error)
	GetCompetitions(ctx context.Context) ([]model.CompetitionWithCategory, error)
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

func (r *MySQLCompetitionRepository) GetCompetitionByCategoryID(ctx context.Context, categoryID uint64) ([]model.Competition, error) {
	var competitions []model.Competition
	if err := r.db.WithContext(ctx).Where("category_id = ?", categoryID).Find(&competitions).Error; err != nil {
		return nil, err
	}
	return competitions, nil
}

func (r *MySQLCompetitionRepository) GetCompetitions(ctx context.Context) ([]model.CompetitionWithCategory, error) {
	var comps []model.CompetitionWithCategory
	if err := r.db.WithContext(ctx).Table("competitions").Select("competitions.*, categories.name as category_name").Joins("JOIN categories ON competitions.category_id = categories.id").Find(&comps).Error; err != nil {
		return nil, err
	}

	return comps, nil
}
