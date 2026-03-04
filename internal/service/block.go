package service

import (
	"campushelphub/internal/errors"
	"campushelphub/internal/repository"
	"campushelphub/model"
	"context"
)

type BlockService struct {
	errs *errors.Error
	repo repository.BlockRepository
}

func NewBlockService(errs *errors.Error, repo repository.BlockRepository) *BlockService {
	return &BlockService{
		errs: errs,
		repo: repo,
	}
}

func (s *BlockService) CreateBlock(ctx context.Context, block *model.Block) error {
	return s.repo.CreateBlock(ctx, block)
}

func (s *BlockService) GetAllBlock(ctx context.Context) ([]model.Block, error) {
	return s.repo.GetAllBlock(ctx)
}
