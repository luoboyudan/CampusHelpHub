package service

import (
	"campushelphub/internal/errors"
	"campushelphub/internal/repository"
	"campushelphub/model"
	"context"
)

type CategoryService struct {
	errs *errors.Error
	repo repository.CategoryRepository
}

func NewCategoryService(errs *errors.Error, repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		errs: errs,
		repo: repo,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, category *model.CreateCategoryRequest) error {
	c := &model.Category{
		Name:        category.Name,
		Description: category.Description,
	}
	return s.repo.CreateCategory(ctx, c)
}

func (s *CategoryService) GetAllCategory(ctx context.Context) ([]model.Category, error) {
	return s.repo.GetAllCategory(ctx)
}
