package errors

const (
	ErrAuth                = "ErrAuth"
	ErrTokenGenerate       = "ErrTokenGenerate"
	ErrUserRegisterRequest = "ErrUserRegisterRequest"
	ErrWechatLoginSession  = "ErrWechatLoginSession"
	ErrUserRegisterCreate  = "ErrUserRegisterCreate"
)

var ErrorMsgTemplates = map[string][2]string{
	ErrUserRegisterRequest: {
		"请求参数错误",
		"请求参数错误: %s",
	},
	ErrWechatLoginSession: {
		"登录失败",
		"登录失败: %s",
	},
	ErrUserRegisterCreate: {
		"创建用户失败",
		"创建用户失败: %s",
	},
	ErrAuth: {
		"认证失败",
		"认证失败: %s",
	},
}
