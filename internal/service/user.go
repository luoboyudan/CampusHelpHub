package service

import (
	"campushelphub/internal/common"
	"campushelphub/internal/repository"
	"campushelphub/internal/service/wechat"
	"campushelphub/model"
	"context"
)

type UserService struct {
	userRepo repository.UserRepository
	IDGen    common.IDgenarator
}

func NewUserService(userRepo repository.UserRepository, idGen common.IDgenarator) *UserService {
	return &UserService{userRepo: userRepo, IDGen: idGen}
}

func (s *UserService) Create(ctx context.Context, req *model.CreateUserRequest) error {
	openID, err := wechat.Login(req.Code)
	if err != nil {
		return err
	}
	user := &model.User{
		ID:       s.IDGen.GenerateID(),
		OpenID:   openID.OpenID,
		Username: req.Username,
		Avatar:   req.Avatar,
		Bio:      req.Bio,
		School:   req.School,
	}
	return s.userRepo.Create(ctx, user)
}
