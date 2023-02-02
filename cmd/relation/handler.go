/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-29 21:58:00
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 18:27:56
 * @FilePath: /ByteCamp/cmd/relation/handler.go
 * @Description: relation微服务handler
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"context"
	"douyin/cmd/relation/pack"
	"douyin/cmd/relation/service"
	relation "douyin/kitex_gen/relation"
	"douyin/pkg/errno"
)

// RelationSrvImpl implements the last service interface defined in the IDL.
type RelationSrvImpl struct{}

var (
	relationService = service.NewRelationService(context.Background())
)

// RelationAction 登录用户对其他用户进行关注或取消关注。
func (s *RelationSrvImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	//1、入参校验
	//2、调用service层，完成关注或取消关注
	err = relationService.RelationAction(req)
	if err != nil {
		resp = pack.BuildRelationActionResponse(err)
		return resp, nil
	}
	//4、返回结果
	return pack.BuildRelationActionResponse(err), nil
}

// 登录用户关注的所有用户列表。
func (s *RelationSrvImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	//1、入参校验
	if req.UserId <= 0 {
		resp = pack.BuildRelationFollowingListResp(nil, errno.ErrBind)
		return resp, nil
	}
	//2、调用service
	users, err := relationService.FollowList(req)
	if err != nil {
		resp = pack.BuildRelationFollowingListResp(nil, err)
		return resp, nil
	}
	resp = pack.BuildRelationFollowingListResp(users, err)
	return resp, nil
}

// 所有关注登录用户的粉丝列表。
func (s *RelationSrvImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	//1、入参校验
	if req.UserId <= 0 {
		resp = pack.BuildRelationFollowerListResp(nil, errno.ErrBind)
		return resp, nil
	}
	//2、调用service
	users, err := relationService.FollowerList(req)
	if err != nil {
		resp = pack.BuildRelationFollowerListResp(nil, err)
		return resp, nil
	}
	resp = pack.BuildRelationFollowerListResp(users, err)
	return resp, nil
}

// RelationFriendList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	if req.UserId <= 0 {
		resp = pack.BuildRelationFriendListResp(nil, errno.ErrBind)
		return resp, nil
	}
	users, err := relationService.FriendList(req)
	if err != nil {
		resp = pack.BuildRelationFriendListResp(nil, err)
		return resp, nil
	}
	resp = pack.BuildRelationFriendListResp(users, err)
	return resp, nil
}
