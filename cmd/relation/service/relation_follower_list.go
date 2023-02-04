/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-04 00:11:19
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-04 13:23:50
 * @FilePath: /ByteCamp/cmd/relation/service/relation_follower_list.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import (
	"context"
	"douyin/cmd/relation/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
	"strconv"
)

//根据req获取RPC所需的粉丝user列表
func (service RelationService) FollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	//1、根据userId获取该user的所有follow列表
	fans, err := db.GetFansByUserId(int(req.UserId))
	if err != nil {
		return nil, err
	}
	fansUsers, err := pack.GetUsersByIds(fans)
	if err != nil {
		return nil, err
	}
	return fansUsers, nil
}

func (service RelationService) FollowerListByRedis(req *relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	//1、查看follower redis中是否有对应的key,若没有，则从mysql中获取到redis中
	ctx := context.Background()
	userIdStr := strconv.Itoa(int(req.UserId))
	cnt, err := db.FollowersRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		return nil, err
	}
	if cnt == 0 {
		loadFollowersListToRedis(ctx, req.UserId)
	} else {
		//更新过期时间
		db.FollowersRedis.Expire(ctx, userIdStr, db.ExpireTime)
	}
	//2、从redis中拿到所有的followerId
	ids, err := getFollowersListFromRedis(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	//3、根据followingId获取user
	followerUsers, err := pack.GetUsersByIds(ids)
	if err != nil {
		return nil, err
	}
	//4、返回user
	return followerUsers, nil
}
