package service

import (
	"campushelphub/internal/errors"
	"campushelphub/internal/repository"
	"campushelphub/model"
	"context"
)

type CompetitionService struct {
	errs *errors.Error
	repo repository.CompetitionRepository
}

func NewCompetitionService(errs *errors.Error, repo repository.CompetitionRepository) *CompetitionService {
	return &CompetitionService{
		errs: errs,
		repo: repo,
	}
}

func (s *CompetitionService) CreateCompetition(ctx context.Context, competition *model.CreateCompetitionRequest) error {
	return s.repo.CreateCompetition(ctx, &competition.Competition)
}

func (s *CompetitionService) GetCompetitionByBlockID(ctx context.Context, blockID uint64) ([]model.Competition, error) {
	return s.repo.GetCompetitionByBlockID(ctx, blockID)
}

func (s *CompetitionService) GetCompetitions(ctx context.Context) ([]model.GetCompetitionResponse, error) {
	comps, err := s.repo.GetCompetitions(ctx)
	if err != nil {
		return nil, err
	}
	BlockMap := make(map[uint]*model.GetCompetitionResponse)
	for _, comp := range comps {
		if BlockMap[comp.BlockID] == nil {
			BlockMap[comp.BlockID] = &model.GetCompetitionResponse{
				BlockID:      comp.BlockID,
				BlockName:    comp.BlockName,
				Competitions: []model.CompetitionNoBlock{},
			}
		} else {
			BlockMap[comp.BlockID].Competitions = append(BlockMap[comp.BlockID].Competitions, model.CompetitionNoBlock{
				ID:         comp.ID,
				Title:      comp.Title,
				EnrollTime: comp.EnrollTime,
				StartTime:  comp.StartTime,
			})
		}
	}
	var res []model.GetCompetitionResponse
	for _, block := range BlockMap {
		res = append(res, *block)
	}
	return res, nil
}
