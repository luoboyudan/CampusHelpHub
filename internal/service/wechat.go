package service

import (
	"campushelphub/internal/config"
	"campushelphub/internal/errors"
	"campushelphub/model"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type WechatService struct {
	config *config.Config
	errs   *errors.Error
}

func NewWechatService(config *config.Config, errs *errors.Error) *WechatService {
	return &WechatService{config: config, errs: errs}
}

func (s *WechatService) Login(code string) (*model.SessionResponse, *errors.Error) {
	params := map[string][]string{
		"appid":      {s.config.Wechat.AppID},
		"secret":     {s.config.Wechat.AppSecret},
		"js_code":    {code},
		"grant_type": {"authorization_code"},
	}
	fullURL := s.config.Wechat.Code2SessionURL + "?" + url.Values(params).Encode()
	resp, err := http.Get(fullURL)
	if err != nil {
		newError := s.errs.NewError(errors.ErrWechatLoginSession, http.StatusInternalServerError, err)
		return nil, newError
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		newError := s.errs.NewError(errors.ErrWechatLoginSession, http.StatusInternalServerError, err)
		return nil, newError
	}
	var sessionResp model.SessionResponse
	err = json.Unmarshal(body, &sessionResp)
	if err != nil {
		newError := s.errs.NewError(errors.ErrWechatLoginSession, http.StatusInternalServerError, err)
		return nil, newError
	}
	if sessionResp.ErrCode != 0 {
		newError := s.errs.NewError(errors.ErrWechatLoginSession, http.StatusInternalServerError, nil)
		return nil, newError
	}
	return &sessionResp, nil
}
