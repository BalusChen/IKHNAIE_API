package constant

const (
	// Basic status code and status message
	StatusCode_OK                    = 200
	StatusCode_Base                  = 50000
	StatusCode_MethodONotImplemented = StatusCode_Base + 1
	StatusCode_InvalidParams         = StatusCode_Base + 2
	StatusMsg_OK                     = "成功"
	StatusMsg_InvalidParams          = "参数错误"
	StatusMsg_ServerInternalError    = "服务器内部错误"
	StatusMsg_MethodNotImplemented   = "该接口尚未实现哦"

	// User-related status code and status message
	StatusCode_UserBase      = StatusCode_Base + 100
	StatusCode_UserNotFound  = StatusCode_UserBase + 1
	StatusCode_UserExist     = StatusCode_UserBase + 2
	StatusCode_WrongPassword = StatusCode_UserBase + 3
	StatusCode_InactiveUser  = StatusCode_UserBase + 4
	StatusMsg_LoginOK        = "登陆成功"
	StatusMsg_UserNotFound   = "该用户不存在"
	StatusMsg_WrongPassword  = "密码错误"
	StatusMsg_RegisterOK     = "注册成功"
	StatusMsg_UserExist      = "该用户已存在"
	StatusMsg_InactiveUser   = "该用户尚未被准入，请联系管理员"

	// Transaction-related status code and status message
	StatusCode_TxBase              = StatusCode_Base + 200
	StatusCode_CallBlockChainError = StatusCode_TxBase + 1
	StatusMsg_CallBalockChainError = "调用区块链服务失败"
)
