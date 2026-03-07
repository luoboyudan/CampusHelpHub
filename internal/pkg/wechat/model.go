package wechat

type SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// template_id 模板ID
const WechatTemplateIDMessageNotice = "default1"
const WechatTemplateIDCompetitionReminder = "default2"

// 模板-跳转URL映射
var WechatJumpURLMap = map[string]string{
	WechatTemplateIDMessageNotice:       "https://campushelphub.com/message",
	WechatTemplateIDCompetitionReminder: "https://campushelphub.com/competition",
}

type WechatReminderData struct {
	Value string `json:"value"`
}
type WechatReminder struct {
	Touser     string                        `json:"touser"`
	TemplateID string                        `json:"template_id"`
	Page       string                        `json:"page"`
	Data       map[string]WechatReminderData `json:"data"`
}
