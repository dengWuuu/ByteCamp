/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 14:46:35
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 18:30:04
 * @FilePath: /ByteCamp/cmd/relation/pack/relation.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package pack

import (
	"context"
	userpack "douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
)

//根据follows列表，获取所有关注用户的rpc格式信息
func GetFollowingByFollows(follows []*db.Follow) ([]*user.User, error) {
	//1、根据follows中的follow_id字段，查询db.User
	ids := make([]int64, 0)
	for _, follow := range follows {
		ids = append(ids, int64(follow.FollowId))
	}
	dbusers, err := db.GetUsersByIds(ids)
	if err != nil {
		return nil, err
	}
	users, err := userpack.Users(context.Background(), dbusers, 0)
	if err != nil {
		return nil, err
	}
	return users, nil
}

//根据follows列表，获取所有粉丝用户的rpc格式信息
func GetFansByFollows(follows []*db.Follow) ([]*user.User, error) {
	//1、根据follows中的user_id字段，查询db.User
	ids := make([]int64, 0)
	for _, follow := range follows {
		ids = append(ids, int64(follow.UserId))
	}
	dbusers, err := db.GetUsersByIds(ids)
	if err != nil {
		return nil, err
	}
	users, err := userpack.Users(context.Background(), dbusers, 0)
	if err != nil {
		return nil, err
	}
	return users, nil
}
