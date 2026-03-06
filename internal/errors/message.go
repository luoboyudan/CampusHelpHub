package errors

const (
	ErrAuth                          = "ErrAuth"
	ErrTokenGenerate                 = "ErrTokenGenerate"
	ErrUserRegisterRequest           = "ErrUserRegisterRequest"
	ErrWechatLoginSession            = "ErrWechatLoginSession"
	ErrUserRegisterCreate            = "ErrUserRegisterCreate"
	ErrUserRegisterGetByWechatOpenID = "ErrUserRegisterGetByWechatOpenID"
	ErrUserVerifyRequest             = "ErrUserVerifyRequest"
	ErrUserVerifyDB                  = "ErrUserVerifyDB"
	ErrChromeOpen                    = "ErrChromeOpen"
	ErrChromeOpenURL                 = "ErrChromeOpenURL"
	ErrChromeInteraction             = "ErrChromeInteraction"
	ErrChromeVerifyFailed            = "ErrChromeVerifyFailed"
	ErrGetPublicKeyRequest           = "ErrGetPublicKeyRequest"
	ErrGetPublicKey                  = "ErrGetPublicKey"
	ErrUserCheckRequest              = "ErrUserCheckRequest"
	ErrUserLoginRequest              = "ErrUserLoginRequest"
	ErrCreateCompetitionRequest      = "ErrCreateCompetitionRequest"
	ErrCreateCompetitionDB           = "ErrCreateCompetitionDB"
	ErrCreateCategoryRequest         = "ErrCreateCategoryRequest"
	ErrGetAllCategoryDB              = "ErrGetAllCategoryDB"
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
	ErrUserRegisterGetByWechatOpenID: {
		Msg:    "通过微信OpenID查询用户失败",
		Detail: "通过微信OpenID查询用户失败(%s)",
	},
	ErrAuth: {
		Msg:    "认证失败",
		Detail: "token认证失败(%s)",
	},
	ErrTokenGenerate: {
		Msg:    "生成token失败",
		Detail: "生成token失败(%s)",
	},
	ErrUserVerifyRequest: {
		Msg:    "验证用户请求参数错误",
		Detail: "验证用户请求参数错误(%s)",
	},
	ErrUserVerifyDB: {
		Msg:    "验证用户数据库错误",
		Detail: "验证用户数据库错误(%s)",
	},
	ErrChromeOpen: {
		Msg:    "打开浏览器会话失败",
		Detail: "打开浏览器会话失败(%s)",
	},
	ErrChromeOpenURL: {
		Msg:    "打开认证页面失败",
		Detail: "打开认证页面失败(%s)",
	},
	ErrChromeInteraction: {
		Msg:    "浏览器交互失败",
		Detail: "浏览器交互失败(%s)",
	},
	ErrChromeVerifyFailed: {
		Msg:    "认证失败",
		Detail: "认证失败(%s)",
	},
	ErrGetPublicKeyRequest: {
		Msg:    "获取公钥请求参数错误",
		Detail: "获取公钥请求参数错误(%s)",
	},
	ErrGetPublicKey: {
		Msg:    "获取公钥失败",
		Detail: "获取公钥失败(%s)",
	},
	ErrUserCheckRequest: {
		Msg:    "检查用户请求参数错误",
		Detail: "检查用户请求参数错误(%s)",
	},
	ErrUserLoginRequest: {
		Msg:    "登录用户请求参数错误",
		Detail: "登录用户请求参数错误(%s)",
	},
	ErrCreateCompetitionRequest: {
		Msg:    "创建竞赛请求参数错误",
		Detail: "创建竞赛请求参数错误(%s)",
	},
	ErrCreateCompetitionDB: {
		Msg:    "创建竞赛失败",
		Detail: "创建竞赛失败(%s)",
	},
	ErrCreateCategoryRequest: {
		Msg:    "创建分类请求参数错误",
		Detail: "创建分类请求参数错误(%s)",
	},
	ErrGetAllCategoryDB: {
		Msg:    "获取所有分类失败",
		Detail: "获取所有分类失败(%s)",
	},
}
