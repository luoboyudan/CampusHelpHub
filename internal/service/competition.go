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

func (s *CompetitionService) GetCompetitionByCategoryID(ctx context.Context, categoryID uint64) ([]model.Competition, error) {
	return s.repo.GetCompetitionByCategoryID(ctx, categoryID)
}

func (s *CompetitionService) GetCompetitions(ctx context.Context) ([]model.GetCompetitionResponse, error) {
	comps, err := s.repo.GetCompetitions(ctx)
	if err != nil {
		return nil, err
	}
	CategoryMap := make(map[uint]*model.GetCompetitionResponse)
	for _, comp := range comps {
		if CategoryMap[comp.CategoryID] == nil {
			CategoryMap[comp.CategoryID] = &model.GetCompetitionResponse{
				CategoryID:   comp.CategoryID,
				CategoryName: comp.CategoryName,
				Competitions: []model.CompetitionNoCategory{},
			}
		} else {
			CategoryMap[comp.CategoryID].Competitions = append(CategoryMap[comp.CategoryID].Competitions, model.CompetitionNoCategory{
				ID:         comp.ID,
				Title:      comp.Title,
				EnrollTime: comp.EnrollTime,
				StartTime:  comp.StartTime,
			})
		}
	}
	var res []model.GetCompetitionResponse
	for _, category := range CategoryMap {
		res = append(res, *category)
	}
	return res, nil
}
