package errors

const (
	ErrAuth                = "ErrAuth"
	ErrTokenGenerate       = "ErrTokenGenerate"
	ErrUserRegisterRequest = "ErrUserRegisterRequest"
	ErrWechatLoginSession  = "ErrWechatLoginSession"
	ErrUserRegisterCreate  = "ErrUserRegisterCreate"
)

type ErrorMsgTemplate struct {
	Msg    string
	Detail string
}

var ErrorMsgTemplates = map[string]ErrorMsgTemplate{
	ErrUserRegisterRequest: {
		Msg:    "请求参数错误",
		Detail: "请求参数错误(%s)",
	},
	ErrWechatLoginSession: {
		Msg:    "登录失败",
		Detail: "微信api接口请求失败(%s)",
	},
	ErrUserRegisterCreate: {
		Msg:    "创建用户失败",
		Detail: "数据库创建用户失败(%s)",
	},
	ErrAuth: {
		Msg:    "认证失败",
		Detail: "token认证失败(%s)",
	},
}
