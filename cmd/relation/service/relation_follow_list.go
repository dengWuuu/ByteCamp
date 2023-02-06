/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 16:41:53
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-06 14:01:40
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
	"strconv"
)

// 根据req获取RPC所需的user列表
func (service RelationService) FollowList(req *relation.DouyinRelationFollowListRequest) ([]*user.User, error) {
	//1、根据userId获取该user的所有follow列表
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
	//1、查看following redis中是否有对应的key,若没有，则从mysql中获取到redis中
	ctx := context.Background()
	userIdStr := strconv.Itoa(int(req.UserId))
	cnt, err := db.FollowingRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		return nil, err
	}
	if cnt == 0 {
		loadFollowingListToRedis(ctx, req.UserId)
	} else {
		//更新过期时间
		db.FollowingRedis.Expire(ctx, userIdStr, db.ExpireTime)
	}
	//2、从redis中拿到所有的followingId
	ids, err := getFollowingListFromRedis(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	//3、根据followingId获取user，先从redis中获取user,若redis中没有，则从mysql中获取
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
	//从redis中获取user
	dbUsers := redis.GetUsersFromRedis(ctx, uids)
	if dbUsers == nil {
		//从mysql中获取user
		followingUsers, err = pack.GetUsersByIds(ids)
		if err != nil {
			return nil, err
		}
	} else {
		//否则直接pack为RPC所需的user
		followingUsers = make([]*user.User, len(dbUsers))
		for i, dbUser := range dbUsers {
			followingUsers[i], err = userpack.User(ctx, dbUser)
			if err != nil {
				return nil, err
			}
		}
	}

	//4、返回user
	return followingUsers, nil
}

//判断一个用户是否关注了另一个用户
func IsFollowing(ctx context.Context, userId, otherId int64) (bool, error) {
	//1.查看redis中是否有缓存
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FollowingRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		return false, err
	}
	if cnt == 0 {
		loadFollowingListToRedis(ctx, userId)
	}
	//2.判断redis缓存中是否有otherId
	otherIdStr := strconv.Itoa(int(otherId))
	exists, err := db.FollowingRedis.SIsMember(ctx, userIdStr, otherIdStr).Result()
	if err != nil {
		return false, err
	}
	return exists, nil
}
