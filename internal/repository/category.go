package repository

import (
	"campushelphub/model"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) error
	GetAllCategory(ctx context.Context) ([]model.Category, error)
}

type MySQLCategoryRepository struct {
	db *gorm.DB
}

func NewMySQLCategoryRepository(db *gorm.DB) CategoryRepository {
	return &MySQLCategoryRepository{
		db: db,
	}
}

func (r *MySQLCategoryRepository) CreateCategory(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySQLCategoryRepository) GetAllCategory(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
