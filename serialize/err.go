package serialize

const (

	// msg
	MsgnicknameOccupy  	= "用户昵称已存在"
	MsgPasswordOccupy  	= "用户密码不一致"
	MsgDBErr			= "数据库操作失败"
	MsgQueryErr			= "查询操作失败"
	MsgParamErr			= "参数错误"
	MsgAccountErr		= "账号或者密码错误"
	MsgEncryptErr		= "密码加密错误"
	MsgCheckLogin		= "未登录"


	// code
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403

	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002

	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	// CodeQueryErr 账号密码错误
	CodeQueryErr = 40002


	// error
	Errdatabase 		= "数据库出错"

)