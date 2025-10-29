package service

import (
	"campushelphub/internal/common"
	"campushelphub/internal/errors"
	"campushelphub/internal/repository"
	"campushelphub/model"
	"context"
	"net/http"
)

type UserService struct {
	userRepo repository.UserRepository
	errs     *errors.Error
	IDGen    common.IDgenarator
}

func NewUserService(userRepo repository.UserRepository, idGen common.IDgenarator) *UserService {
	return &UserService{userRepo: userRepo, IDGen: idGen}
}

func (s *UserService) Create(ctx context.Context, req *model.CreateUserRequest, sessionResp *model.SessionResponse) *errors.Error {
	user := &model.User{
		ID:       s.IDGen.GenerateID(),
		OpenID:   sessionResp.OpenID,
		Username: req.Username,
		Avatar:   req.Avatar,
		Bio:      req.Bio,
		School:   req.School,
	}
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return s.errs.NewError("创建用户失败", err.Error(), http.StatusInternalServerError, err)
	}
	return nil
}
