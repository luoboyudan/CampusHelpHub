package errors

const (
	ErrAuth                          = "ErrAuth"
	ErrBadRequest                    = "ErrBadRequest"
	ErrTokenGenerate                 = "ErrTokenGenerate"
	ErrWechatLoginSession            = "ErrWechatLoginSession"
	ErrUserRegisterCreate            = "ErrUserRegisterCreate"
	ErrUserRegisterGetByWechatOpenID = "ErrUserRegisterGetByWechatOpenID"
	ErrUserCheckDB                   = "ErrUserCheckDB"
	ErrUserVerifyDB                  = "ErrUserVerifyDB"
	ErrChromeOpen                    = "ErrChromeOpen"
	ErrChromeOpenURL                 = "ErrChromeOpenURL"
	ErrChromeInteraction             = "ErrChromeInteraction"
	ErrChromeVerifyFailed            = "ErrChromeVerifyFailed"
	ErrGetPublicKey                  = "ErrGetPublicKey"
	ErrCreateCompetitionDB           = "ErrCreateCompetitionDB"
	ErrCreateCategoryDB              = "ErrCreateCategoryDB"
	ErrGetAllCategoryDB              = "ErrGetAllCategoryDB"
	ErrParamEmpty                    = "ErrParamEmpty"
)

type ErrorMsgTemplate struct {
	Msg    string
	Detail string
}

var ErrorMsgTemplates = map[string]ErrorMsgTemplate{
	ErrBadRequest: {
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
	ErrGetPublicKey: {
		Msg:    "获取公钥失败",
		Detail: "获取公钥失败(%s)",
	},
	ErrCreateCompetitionDB: {
		Msg:    "创建竞赛失败",
		Detail: "创建竞赛失败(%s)",
	},
	ErrGetAllCategoryDB: {
		Msg:    "获取所有分类失败",
		Detail: "获取所有分类失败(%s)",
	},
	ErrUserCheckDB: {
		Msg:    "检查用户数据库错误",
		Detail: "检查用户数据库错误(%s)",
	},
	ErrParamEmpty: {
		Msg:    "参数为空",
		Detail: "参数为空(%s)",
	},
}
