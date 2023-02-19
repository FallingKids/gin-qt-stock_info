package constant

type Message string

const (
	SUCCESS_MESSAGE         Message = "success"
	PARAMS_VALIDATE_ERROR   Message = "参数校验错误"
	READALL_ERROR           Message = "数据读取错误"
	JSON_UNMARSHAL_ERROR    Message = "数据解析错误"
	UNKNOWN_ERROR           Message = "系统未知错误"
	USER_IS_NOT_EXIST_ERROR Message = "用户不存在"
	TOKEN_INVALID_ERROR     Message = "登录信息已经过期"
)
