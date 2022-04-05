package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code= 1000... 认证模块错误
	ERROR_USERNAME_USED         = 1001
	ERROR_USERNAME_WRONG        = 1002
	ERROR_PASSWORD_WRONG        = 1003
	ERROR_USER_NOT_EXIST        = 1004
	ERROR_TOKEN_EXPIRED         = 1005
	ERROR_TOKEN_WRONG           = 1006
	ERROR_GENERATE_TOKEN_FAILED = 1007

	// code= 2000... 信息模块错误
	//ERROR_DOMAIN_WRONG    = 2001
	ERROR_GET_INFO_FAILED = 2002
)

var codeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "用户名已存在",
	ERROR_USERNAME_WRONG: "用户名错误",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户名错误或不存在",

	ERROR_TOKEN_EXPIRED: "TOKEN已过期",
	ERROR_TOKEN_WRONG:   "TOKEN不正确",

	//ERROR_DOMAIN_WRONG:    "域名错误",
	ERROR_GET_INFO_FAILED: "获取信息失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
