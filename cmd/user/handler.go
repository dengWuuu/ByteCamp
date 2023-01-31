package main

import (
	"context"
	"douyin/cmd/user/service"
	user "douyin/kitex_gen/user"
	"douyin/pack"
	"douyin/pkg/errno"
)

// UserSrvImpl implements the last service interface defined in the IDL.
type UserSrvImpl struct{}

// Register implements the UserSrvImpl interface.
func (s *UserSrvImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	if len(req.Username) == 0 || len(req.Password) == 0 {

		resp = pack.BuildUserRegisterResp(errno.ErrBind)
		return resp, nil
	}
	err = service.NewRegisterService(ctx).Register(req)
	if err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return resp, nil
	}

	//包装成功响应
	resp = pack.BuildUserRegisterResp(errno.Success)
	return resp, nil
}

// Login implements the UserSrvImpl interface.
func (s *UserSrvImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the UserSrvImpl interface.
func (s *UserSrvImpl) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}
