/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 00:49:20
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 18:17:58
 * @FilePath: /ByteCamp/cmd/relation/pack/relation_resp.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */

package pack

import (
	"errors"

	relation "douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
)

// 构造关注或取消关注RPC response
// 根据error，判断其属于那种errno.ErrNo,调用GetRelationActionResp返回对应的relation.DouyinRelationActionResponse
func BuildRelationActionResponse(err error) *relation.DouyinRelationActionResponse {
	if err == nil {
		return getRelationActionResp(errno.Success)
	} else {
		e := errno.ErrNo{}
		if errors.As(err, &e) {
			return getRelationActionResp(e)
		}

		s := errno.ErrUnknown.WithMessage(err.Error())
		return getRelationActionResp(s)
	}
}

// 根据传入的errno.ErrNo，返回对应的relation.DouyinRelationActionResponse
func getRelationActionResp(err errno.ErrNo) *relation.DouyinRelationActionResponse {
	return &relation.DouyinRelationActionResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}

// 构造关注列表RPC response
func BuildRelationFollowingListResp(users []*user.User, err error) *relation.DouyinRelationFollowListResponse {
	if err == nil {
		return getRelationFollowingListResp(users, errno.Success)
	} else {
		e := errno.ErrNo{}
		if errors.As(err, &e) {
			return getRelationFollowingListResp(users, e)
		}

		s := errno.ErrUnknown.WithMessage(err.Error())
		return getRelationFollowingListResp(users, s)
	}
}

func getRelationFollowingListResp(users []*user.User, err errno.ErrNo) *relation.DouyinRelationFollowListResponse {
	return &relation.DouyinRelationFollowListResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
		UserList:   users,
	}
}

// 构造粉丝列表RPC response
func BuildRelationFollowerListResp(users []*user.User, err error) *relation.DouyinRelationFollowerListResponse {
	if err == nil {
		return getRelationFollowerListResp(users, errno.Success)
	} else {
		e := errno.ErrNo{}
		if errors.As(err, &e) {
			return getRelationFollowerListResp(users, e)
		}

		s := errno.ErrUnknown.WithMessage(err.Error())
		return getRelationFollowerListResp(users, s)
	}
}

func getRelationFollowerListResp(users []*user.User, err errno.ErrNo) *relation.DouyinRelationFollowerListResponse {
	return &relation.DouyinRelationFollowerListResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
		UserList:   users,
	}
}

// 构造朋友列表RPC response
func BuildRelationFriendListResp(users []*user.User, err error) *relation.DouyinRelationFriendListResponse {
	if err == nil {
		return getRelationFriendListResp(users, errno.Success)
	} else {
		e := errno.ErrNo{}
		if errors.As(err, &e) {
			return getRelationFriendListResp(users, e)
		}

		s := errno.ErrUnknown.WithMessage(err.Error())
		return getRelationFriendListResp(users, s)
	}
}

func getRelationFriendListResp(users []*user.User, err errno.ErrNo) *relation.DouyinRelationFriendListResponse {
	return &relation.DouyinRelationFriendListResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
		UserList:   users,
	}
}
