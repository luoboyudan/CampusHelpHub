package repository

import (
	"campushelphub/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByWechatOpenID(ctx context.Context, wechatOpenID string) (*model.User, error)
	Verify(ctx context.Context, userID uint64) error
}

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) UserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *MySQLUserRepository) GetByWechatOpenID(ctx context.Context, wechatOpenID string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("wechat_open_id = ?", wechatOpenID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) Verify(ctx context.Context, userID uint64) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).Update("auth", true).Error
}
