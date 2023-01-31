package pack

import (
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"errors"
)

// BuildUserRegisterResp build getUserRegisterResp from error
func BuildUserRegisterResp(err error) *user.DouyinUserRegisterResponse {
	if err == nil {
		return getUserRegisterResp(errno.Success)
	}
	//如果是定义的错误则打印
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getUserRegisterResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return getUserRegisterResp(s)
}

func getUserRegisterResp(err errno.ErrNo) *user.DouyinUserRegisterResponse {
	return &user.DouyinUserRegisterResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func BuildGetUserResp(err error) *user.DouyinUserResponse {
	if err == nil {
		return getGetUserResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getGetUserResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return getGetUserResp(s)
}

func getGetUserResp(err errno.ErrNo) *user.DouyinUserResponse {
	return &user.DouyinUserResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
