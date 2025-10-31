package auth

import (
	"time"

	"campushelphub/internal/errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager struct {
	SecretKey  []byte
	ExpireTime int
	Error      *errors.Error
}

func NewTokenManager(secretKey string, expireTime int) *TokenManager {
	return &TokenManager{
		SecretKey:  []byte(secretKey),
		ExpireTime: expireTime,
	}
}

func (tm *TokenManager) GenerateToken(userID int64) (string, error) {
	// 生成token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(tm.ExpireTime) * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(tm.SecretKey)
	if err != nil {
		return "", tm.Error.NewError("生成token失败", err.Error(), http.StatusInternalServerError, err)
	}
	return tokenStr, nil
}

func (tm *TokenManager) VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	// 验证token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return tm.SecretKey, nil
	})
	if err != nil {
		return nil, tm.Error.NewError("验证token失败", err.Error(), http.StatusUnauthorized, err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, tm.Error.NewError("无效的token claims", "invalid token claims", http.StatusUnauthorized, err)
	}
	return claims, nil
}
