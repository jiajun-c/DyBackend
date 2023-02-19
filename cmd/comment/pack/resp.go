package pack

import (
	"errors"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/base"
)

//BuildCommentBaseResp build comment baseResp from error
func BuildCommentBaseResp(err error) *base.BaseResponse {
	if err == nil {
		return commentbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return commentbaseResp(s)
}

func commentbaseResp(err errno.ErrNo) *base.BaseResponse {
	return &base.BaseResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
