package service

import (
	"campushelphub/internal/common/auth"
	"campushelphub/internal/common/snowflake"
	"campushelphub/internal/errors"
	"campushelphub/internal/repository"
	"campushelphub/model"
	"context"
	"net/http"
)

type UserService struct {
	userRepo     repository.UserRepository
	errs         *errors.Error
	IDGen        snowflake.IDgenarator
	TokenManager *auth.TokenManager
}

func NewUserService(userRepo repository.UserRepository, idGen snowflake.IDgenarator, tokenManager *auth.TokenManager) *UserService {
	return &UserService{userRepo: userRepo, IDGen: idGen, TokenManager: tokenManager}
}

func (s *UserService) CheckUser(ctx context.Context, openid string) (bool, *errors.Error) {
	exist, err := s.userRepo.CheckUserExist(ctx, openid)
	if err != nil {
		return false, s.errs.NewError(errors.ErrUserCheckRequest, http.StatusInternalServerError, err)
	}
	return exist, nil
}

func (s *UserService) Login(ctx context.Context, sessionResp *model.SessionResponse) (*model.User, *errors.Error) {
	user, err := s.userRepo.GetByWechatOpenID(ctx, sessionResp.OpenID)
	if err != nil {
		return nil, s.errs.NewError(errors.ErrUserLoginRequest, http.StatusInternalServerError, err)
	}
	return user, nil
}

func (s *UserService) Create(ctx context.Context, req *model.CreateUserRequest, sessionResp *model.SessionResponse) (*model.User, *errors.Error) {
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
		return nil, s.errs.NewError(errors.ErrUserRegisterCreate, http.StatusInternalServerError, err)
	}
	return user, nil
}

func (s *UserService) Verify(ctx context.Context, req *model.VerifyUserRequest) *errors.Error {
	err := s.userRepo.Verify(ctx, req.UserID)
	if err != nil {
		return s.errs.NewError(errors.ErrUserVerifyDB, http.StatusInternalServerError, err)
	}
	return nil
}
