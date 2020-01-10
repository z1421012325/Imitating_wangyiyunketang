package serialize

import "github.com/gin-gonic/gin"


// 通用模板
type Response struct {
	Code int              	`json:"code"`
	Msg  string           	`json:"msg"`
	Data interface{}      	`json:"data"`
	Err  string				`json:"err"`
}


// Res 通用返回
func Res(data interface{}, msg string) *Response {
	res := &Response{
		Data: data,
		Msg:  msg,
	}
	return res
}





















// Err 通用错误处理
func Err(errCode int, msg string, err error) *Response {
	res := &Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Err = err.Error()
	}
	return res
}



// DBErr 数据库操作失败
func DBErr(msg string, err error) *Response {
	if msg == "" {
		msg = MsgDBErr
	}
	return Err(CodeDBError, msg, err)
}

// QueryErr 查询操作失败
func QueryErr(msg string, err error) *Response {
	if msg == "" {
		msg = MsgQueryErr
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) *Response {
	if msg == "" {
		msg = MsgParamErr
	}
	return Err(CodeParamErr, msg, err)
}

// AccountErr 账号或者密码错误
func AccountErr(msg string, err error) *Response {
	if msg == "" {
		msg = MsgAccountErr
	}
	return Err(CodeQueryErr, msg, err)
}

// PswdErr 密码不一致
func PswdErr(msg string, err error) *Response {
	if msg == "" {
		msg = MsgPasswordOccupy
	}
	return Err(CodeQueryErr, msg, err)
}

// EncryptErr 密码加密错误
func EncryptionErr(msg string, err error) *Response {
	if msg == "" {
		msg = MsgEncryptErr
	}
	return Err(CodeEncryptError, msg, err)
}


func TransactionErr(msg string, err error) *Response {
	if msg == "" {
		msg = MsgTransaction
	}
	return Err(CodeTransactionError, msg, err)
}









// CheckLogin 检查登录
func CheckLogin() *Response {
	return &Response{
		Code: CodeCheckLogin,
		Msg:  MsgCheckLogin,
	}
}
