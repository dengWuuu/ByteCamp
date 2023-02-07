/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-04 00:11:19
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-06 15:48:33
 * @FilePath: \ByteCamp\cmd\relation\service\relation_follower_list.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import (
	"context"

	"douyin/cmd/relation/pack"
	userpack "douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
	redis "douyin/pkg/redis"
)

// 根据req获取RPC所需的粉丝user列表
func (service RelationService) FollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	// 1、根据userId获取该user的所有follow列表
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
	// 1、load
	ctx := context.Background()
	loadFollowersListToRedis(ctx, req.UserId)
	// 2、从redis中拿到所有的followerId
	ids, err := getFollowersListFromRedis(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// 3、根据followingId获取user
	var followerUsers []*user.User
	uids := make([]uint, len(ids)-1)
	k := 0
	for _, id := range ids {
		if id == -1 {
			continue
		}
		uids[k] = uint(id)
		k++
	}
	// 从redis中获取user
	dbUsers := redis.GetUsersFromRedis(ctx, uids)
	if dbUsers == nil {
		// 从mysql中获取user
		followerUsers, err = pack.GetUsersByIds(ids)
		if err != nil {
			return nil, err
		}
	} else {
		// 否则直接pack为RPC所需的user
		followerUsers = make([]*user.User, len(dbUsers))
		for i, dbUser := range dbUsers {
			followerUsers[i], err = userpack.User(ctx, dbUser)
			if err != nil {
				return nil, err
			}
		}
	}

	// 4、补全users中的关注总数、粉丝总数、是否关注
	for _, followerUser := range followerUsers {
		// loadFollowersListToRedis(ctx, followerUser.Id)
		// loadFollowingListToRedis(ctx, followerUser.Id)
		// followcnt, err := getFollowingCountFromRedis(ctx, followerUser.Id)
		// if err != nil {
		// 	return nil, err
		// }
		// followerUser.FollowCount = &followcnt

		// followercnt, err := getFollowersCountFromRedis(ctx, followerUser.Id)
		// if err != nil {
		// 	return nil, err
		// }
		// followerUser.FollowerCount = &followercnt
		isFollow, err := redis.IsFollowing(ctx, req.UserId, followerUser.Id)
		if err != nil {
			return nil, err
		}
		followerUser.IsFollow = isFollow
	}
	// 5、返回user
	return followerUsers, nil
}
