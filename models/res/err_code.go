// 自定义错误信息

package res

type ErrorCode int

const (
	SettingsError       ErrorCode = 1001 // 系统错误
	ArgumentError       ErrorCode = 1002 //参数错误
	UnauthorizedError   ErrorCode = 1003 //未授权
	NotFoundError       ErrorCode = 1004 //未找到
	ForbiddenError      ErrorCode = 1005 //禁止访问
	InternalServerError ErrorCode = 1006 //内部服务器错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:       "系统错误",
		ArgumentError:       "参数错误",
		UnauthorizedError:   "未授权",
		NotFoundError:       "未找到",
		ForbiddenError:      "禁止访问",
		InternalServerError: "内部服务器错误",
	}
)
