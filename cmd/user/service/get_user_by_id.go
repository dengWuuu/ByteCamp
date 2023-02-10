package service

import (
	"context"

	"douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) GetUserService {
	return GetUserService{ctx: ctx}
}

func (getUserService GetUserService) GetUserById(req *user.DouyinUserRequest) (*user.User, error) {
	model, err := db.GetUserById(req.UserId)
	if err != nil {
		return nil, err
	}
	rpcUser, err := pack.User(getUserService.ctx, model, req.FromId)
	if err != nil {
		return nil, err
	}
	return rpcUser, nil
}
