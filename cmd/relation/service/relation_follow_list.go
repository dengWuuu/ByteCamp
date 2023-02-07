/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 16:41:53
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-06 15:43:26
 * @FilePath: \ByteCamp\cmd\relation\service\relation_follow_list.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
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
	"context"

	"douyin/cmd/relation/pack"
	userpack "douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
	redis "douyin/pkg/redis"
)

// 根据req获取RPC所需的user列表
func (service RelationService) FollowList(req *relation.DouyinRelationFollowListRequest) ([]*user.User, error) {
	// 1、根据userId获取该user的所有follow列表
	ids, err := db.GetFollowingByUserId(int(req.UserId))
	if err != nil {
		return nil, err
	}
	followingUsers, err := pack.GetUsersByIds(ids)
	if err != nil {
		return nil, err
	}
	return followingUsers, nil
}

func (service RelationService) FollowListByRedis(req *relation.DouyinRelationFollowListRequest) ([]*user.User, error) {
	// 1、load
	ctx := context.Background()
	loadFollowingListToRedis(ctx, req.UserId)
	// 2、从redis中拿到所有的followingId
	ids, err := getFollowingListFromRedis(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// 3、根据followingId获取user，先从redis中获取user,若redis中没有，则从mysql中获取
	var followingUsers []*user.User
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
		followingUsers, err = pack.GetUsersByIds(ids)
		if err != nil {
			return nil, err
		}
	} else {
		// 否则直接pack为RPC所需的user
		followingUsers = make([]*user.User, len(dbUsers))
		for i, dbUser := range dbUsers {
			followingUsers[i], err = userpack.User(ctx, dbUser)
			if err != nil {
				return nil, err
			}
		}
	}
	// 4、补全users中的关注总数、粉丝总数、是否关注
	for _, followingUser := range followingUsers {
		// loadFollowersListToRedis(ctx, followingUser.Id)
		// loadFollowingListToRedis(ctx, followingUser.Id)
		// followcnt, err := getFollowingCountFromRedis(ctx, followingUser.Id)
		// if err != nil {
		// 	return nil, err
		// }
		// followingUser.FollowCount = &followcnt

		// followercnt, err := getFollowersCountFromRedis(ctx, followingUser.Id)
		// if err != nil {
		// 	return nil, err
		// }
		// followingUser.FollowerCount = &followercnt
		isFollow, err := redis.IsFollowing(ctx, req.UserId, followingUser.Id)
		if err != nil {
			return nil, err
		}
		followingUser.IsFollow = isFollow
	}

	// 5、返回user
	return followingUsers, nil
}
