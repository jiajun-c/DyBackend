package app

import (
	"errors"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/base"
)

func BuildBaseResp(err error) *base.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *base.BaseResponse {
	return &base.BaseResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  err.ErrMsg,
	}
}
