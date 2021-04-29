package ecode

const (
	// success
	StatusOK = 0

	// first time use out leave service
	StatusFirstUseLeave = 4

	// error params
	StatusErrorParams = -1
	// internal error
	StatusInternalError = -2
	// unauthorized
	StatusUnauthorized = -3
	// login failure
	StatusLoginFailure = -4

	// unknown status code
	StatusUnknown = -9
)

var errMsg = map[int]string{
	StatusOK:            "操作成功",
	StatusFirstUseLeave: "用户首次使用",
	StatusErrorParams:   "参数错误",
	StatusInternalError: "服务器内部错误",
	StatusUnauthorized:  "您没有权限,请联系管理员",
	StatusLoginFailure:  "登录失效",
	StatusUnknown:       "未知状态码",
}

// GetMsg is a proxy,to prevent errMsg from being modified
func GetMsg(code int) string {
	return errMsg[code]
}
