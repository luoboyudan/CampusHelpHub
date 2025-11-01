package auth

import (
	"time"

	"campushelphub/internal/config"
	"campushelphub/internal/errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager struct {
	SecretKey  []byte
	ExpireTime int
	Error      *errors.Error
}

func NewTokenManager(Config *config.Config) *TokenManager {
	return &TokenManager{
		SecretKey:  []byte(Config.Token.SecretKey),
		ExpireTime: Config.Token.ExpireTime,
	}
}
func (tm *TokenManager) GenerateToken(userID int64) (string, *errors.Error) {
	// 生成token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(tm.ExpireTime) * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(tm.SecretKey)
	if err != nil {
		return "", tm.Error.NewError(errors.ErrTokenGenerate, http.StatusInternalServerError, err)
	}
	return tokenStr, nil
}

func (tm *TokenManager) VerifyToken(tokenStr string) (jwt.MapClaims, *errors.Error) {
	// 验证token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return tm.SecretKey, nil
	})
	if err != nil {
		return nil, tm.Error.NewError(errors.ErrAuth, http.StatusUnauthorized, err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, tm.Error.NewError(errors.ErrAuth, http.StatusUnauthorized, err)
	}
	return claims, nil
}
