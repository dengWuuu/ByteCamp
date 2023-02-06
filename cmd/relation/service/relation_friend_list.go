/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-02 18:43:44
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-06 16:02:57
 * @FilePath: \ByteCamp\cmd\relation\service\relation_friend_list.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-02 18:43:44
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-03 23:09:26
 * @FilePath: /ByteCamp/cmd/relation/service/relation_friend_list.go
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

//根据req获取RPC所需的朋友userId列表
func (service RelationService) FriendList(req *relation.DouyinRelationFriendListRequest) ([]*user.User, error) {
	//1、根据userId获取该user的所有follow列表
	ids, err := db.GetFriendsByUserId(int(req.UserId))
	if err != nil {
		return nil, err
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

func (service RelationService) FriendListByRedis(req *relation.DouyinRelationFriendListRequest) ([]*user.User, error) {
	//1、查看Friend redis中是否有对应的key,若没有，则从mysql中获取到redis中
	ctx := context.Background()
	loadFriendsListToRedis(ctx, req.UserId)
	//2、从redis中拿到所有的FriendId
	ids, err := getFriendsListFromRedis(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	//3、根据followingId获取user
	var FriendUsers []*user.User
	uids := make([]uint, len(ids)-1)
	k := 0
	for _, id := range ids {
		if id == -1 {
			continue
		}
		uids[k] = uint(id)
		k++
	}
	dbUsers := redis.GetUsersFromRedis(ctx, uids)
	if dbUsers == nil {
		//从mysql中获取user
		FriendUsers, err = pack.GetUsersByIds(ids)
		if err != nil {
			return nil, err
		}
	} else {
		//否则直接pack为RPC所需的user
		FriendUsers = make([]*user.User, len(dbUsers))
		for i, dbUser := range dbUsers {
			FriendUsers[i], err = userpack.User(ctx, dbUser)
			if err != nil {
				return nil, err
			}
		}
	}
	if err != nil {
		return nil, err
	}
	//4、补全信息
	for _, friendUser := range FriendUsers {
		// loadFollowersListToRedis(ctx, friendUser.Id)
		// loadFollowingListToRedis(ctx, friendUser.Id)
		// followcnt, err := getFollowingCountFromRedis(ctx, friendUser.Id)
		// if err != nil {
		// 	return nil, err
		// }
		// friendUser.FollowCount = &followcnt

		// followercnt, err := getFollowersCountFromRedis(ctx, friendUser.Id)
		// if err != nil {
		// 	return nil, err
		// }
		// friendUser.FollowerCount = &followercnt
		isFollow, err := redis.IsFollowing(ctx, req.UserId, friendUser.Id)
		if err != nil {
			return nil, err
		}
		friendUser.IsFollow = isFollow
	}
	//5、返回user
	return FriendUsers, nil
}
