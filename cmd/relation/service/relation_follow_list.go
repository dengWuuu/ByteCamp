/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 15:29:24
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 15:41:48
 * @FilePath: /ByteCamp/cmd/relation/service/relation_follow_list.go
 * @Description:调用pack、db，完成对数据库的操作
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import (
	"douyin/cmd/relation/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
)

//根据req获取RPC所需的user列表
func (service RelationService) FollowList(req *relation.DouyinRelationFollowListRequest) ([]*user.User, error) {
	//1、根据userId获取该user的所有follow列表
	followings, err := db.GetFollowingByUserId(int(req.UserId))
	if err != nil {
		return nil, err
	}
	followingUsers, err := pack.GetFollowingByFollows(followings)
	if err != nil {
		return nil, err
	}
	return followingUsers, nil
}
