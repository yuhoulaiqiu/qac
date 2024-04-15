// 自定义错误信息

package res

type ErrorCode int

const (
	SettingsError         ErrorCode = 1001 // 系统错误
	ArgumentError         ErrorCode = 1002 //参数错误
	UnauthorizedError     ErrorCode = 1003 //未授权
	NotFoundError         ErrorCode = 1004 //未找到
	ForbiddenError        ErrorCode = 1005 //禁止访问
	InternalServerError   ErrorCode = 1006 //内部服务器错误
	RedisError            ErrorCode = 1007 //redis错误
	RepeatDiggError       ErrorCode = 1008 //重复点赞
	InvalidOperationError ErrorCode = 1009 //无效操作
	DBError               ErrorCode = 1010 //mysql数据库错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:         "系统错误",
		ArgumentError:         "参数错误",
		UnauthorizedError:     "未授权",
		NotFoundError:         "未找到",
		ForbiddenError:        "禁止访问",
		InternalServerError:   "内部服务器错误",
		RedisError:            "redis错误",
		RepeatDiggError:       "重复点赞",
		InvalidOperationError: "无效操作",
		DBError:               "mysql数据库错误",
	}
)
