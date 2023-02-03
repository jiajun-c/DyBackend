// Package errno
// @Description: 传递的错误信息
package errno

import "fmt"

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyErrCode         = 10003
	AuthorizationFailedErrCode = 10004
	UserRegisterFailedErrCode  = 10005
)

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong param")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyErrCode, "The User already exist")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserRegisterFailedErr  = NewErrNo(UserRegisterFailedErrCode, "Register failed")
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}
