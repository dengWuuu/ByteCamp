/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-03 22:16:48
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-10 14:32:42
 * @FilePath: /ByteCamp/cmd/relation/service/relation_redis.go
 * @Description: relation微服务对redis的操作封装
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import (
	"context"
	"strconv"

	"douyin/dal/db"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// 向redis的following set中添加关注
func addRedisFollowList(ctx context.Context, userId, toUserId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	loadFollowingListToRedis(ctx, userId)
	err := db.FollowingRedis.SAdd(ctx, userIdStr, toUserId).Err()
	if err != nil {
		return err
	}
	err = db.FollowingRedis.Expire(ctx, userIdStr, db.ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// 向redis的follower set中添加粉丝
func addRedisFollowerList(ctx context.Context, userId, toUserId int64) error {
	toUserIdStr := strconv.Itoa(int(toUserId))
	loadFollowersListToRedis(ctx, toUserId)
	err := db.FollowersRedis.SAdd(ctx, toUserIdStr, userId).Err()
	if err != nil {
		return err
	}
	err = db.FollowersRedis.Expire(ctx, toUserIdStr, db.ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// 向redis的friends set中添加朋友
func addRedisFriendsList(ctx context.Context, userId, toUserId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	loadFriendsListToRedis(ctx, userId)
	// 如果redis中不存在该用户的关注列表，那么不对redis进行操作
	// TODO:如果要添加朋友关系，首先必须将toUser的following set添加进来(Done)
	// 1、将toUser的following set添加进来
	toUserIdStr := strconv.Itoa(int(toUserId))
	loadFollowingListToRedis(ctx, toUserId)
	// 2、判断其中是否存在userId
	isMember, err := db.FollowingRedis.SIsMember(ctx, toUserIdStr, userId).Result()
	if err != nil {
		return err
	}
	// 3、如果存在，则说明添加完当前的关注关系后，两者之间存在朋友关系，需要更新当前user和toUser在redis中的friends set
	if isMember {
		err = db.FriendsRedis.SAdd(ctx, userIdStr, toUserId).Err()
		if err != nil {
			return err
		}
		err = db.FriendsRedis.Expire(ctx, userIdStr, db.ExpireTime).Err()
		if err != nil {
			return err
		}
		err = db.FriendsRedis.SAdd(ctx, toUserIdStr, userId).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// 向redis的following set中移除关注
func remRedisFollowList(ctx context.Context, userId, toUserId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	loadFollowingListToRedis(ctx, userId)
	// 如果redis中不存在该用户的关注列表，那么不对redis进行操作
	err := db.FollowingRedis.SRem(ctx, userIdStr, toUserId).Err()
	if err != nil {
		return err
	}
	err = db.FollowingRedis.Expire(ctx, userIdStr, db.ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// 向redis的follower set中移除粉丝
func remRedisFollowerList(ctx context.Context, userId, toUserId int64) error {
	toUserIdStr := strconv.Itoa(int(toUserId))
	loadFollowersListToRedis(ctx, toUserId)
	err := db.FollowersRedis.SRem(ctx, toUserIdStr, userId).Err()
	if err != nil {
		return err
	}
	err = db.FollowersRedis.Expire(ctx, toUserIdStr, db.ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// 向redis的friends set中添加朋友
func remRedisFriendsList(ctx context.Context, userId, toUserId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	loadFriendsListToRedis(ctx, userId)
	err := db.FriendsRedis.SRem(ctx, userIdStr, toUserId).Err()
	if err != nil {
		return err
	}
	err = db.FriendsRedis.Expire(ctx, userIdStr, db.ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从数据库中加载关注列表,并将其存入redis
func loadFollowingListToRedis(ctx context.Context, userId int64) error {
	// 首先判断该用户的缓存是否已经存在，如果存在，则不需要再次加载
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FollowingRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		hlog.Infof("加载关注列表到redis失败")
		return err
	}
	if cnt != 0 {
		db.FollowingRedis.Expire(ctx, strconv.Itoa(int(userId)), db.ExpireTime)
		return nil
	}
	// 从数据库load时需要首先添加一个-1的key，防止读脏
	err = db.FollowingRedis.SAdd(ctx, strconv.Itoa(int(userId)), -1).Err()
	if err != nil {
		return err
	}
	ids, err := db.GetFollowingByUserId(int(userId))
	if err != nil {
		return err
	}
	if len(ids) > 0 {
		var idsStr []string
		for _, id := range ids {
			idsStr = append(idsStr, strconv.Itoa(int(id)))
		}
		err = db.FollowingRedis.SAdd(ctx, strconv.Itoa(int(userId)), idsStr).Err()
		if err != nil {
			return err
		}
	}
	err = db.FollowingRedis.Expire(ctx, strconv.Itoa(int(userId)), db.ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从数据库中加载粉丝列表,并将其存入redis
func loadFollowersListToRedis(ctx context.Context, userId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FollowersRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		hlog.Fatal("加载粉丝列表到redis失败")
		return err
	}
	if cnt != 0 {
		db.FollowersRedis.Expire(ctx, strconv.Itoa(int(userId)), db.ExpireTime)
		return nil
	}
	err = db.FollowersRedis.SAdd(ctx, strconv.Itoa(int(userId)), -1).Err()
	if err != nil {
		return err
	}
	ids, err := db.GetFansByUserId(int(userId))
	if err != nil {
		return err
	}
	if len(ids) > 0 {
		var idsStr []string
		for _, id := range ids {
			idsStr = append(idsStr, strconv.Itoa(int(id)))
		}
		err = db.FollowersRedis.SAdd(ctx, strconv.Itoa(int(userId)), idsStr).Err()
		if err != nil {
			return err
		}
	}
	err = db.FollowersRedis.Expire(ctx, strconv.Itoa(int(userId)), db.ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从数据库中加载朋友列表,并将其存入redis
func loadFriendsListToRedis(ctx context.Context, userId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FriendsRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		hlog.Infof("加载好友列表到redis失败")
		return err
	}
	if cnt != 0 {
		db.FriendsRedis.Expire(ctx, strconv.Itoa(int(userId)), db.ExpireTime)
		return nil
	}
	err = db.FriendsRedis.SAdd(ctx, strconv.Itoa(int(userId)), -1).Err()
	if err != nil {
		return err
	}
	ids, err := db.GetFriendsByUserId(int(userId))
	if err != nil {
		return err
	}
	if len(ids) > 0 {
		var idsStr []string
		for _, id := range ids {
			idsStr = append(idsStr, strconv.Itoa(int(id)))
		}
		err = db.FriendsRedis.SAdd(ctx, strconv.Itoa(int(userId)), idsStr).Err()
		if err != nil {
			return err
		}
	}
	err = db.FriendsRedis.Expire(ctx, strconv.Itoa(int(userId)), db.ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// 从redis中获取关注列表
func getFollowingListFromRedis(ctx context.Context, userId int64) ([]int64, error) {
	ids, err := db.FollowingRedis.SMembers(ctx, strconv.Itoa(int(userId))).Result()
	if err != nil {
		return nil, err
	}
	var res []int64
	for _, id := range ids {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		res = append(res, int64(idInt))
	}
	return res, nil
}

// 从redis中获取粉丝列表
func getFollowersListFromRedis(ctx context.Context, userId int64) ([]int64, error) {
	ids, err := db.FollowersRedis.SMembers(ctx, strconv.Itoa(int(userId))).Result()
	if err != nil {
		return nil, err
	}
	var res []int64
	for _, id := range ids {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		res = append(res, int64(idInt))
	}
	return res, nil
}

// 从redis中获取朋友列表
func getFriendsListFromRedis(ctx context.Context, userId int64) ([]int64, error) {
	ids, err := db.FriendsRedis.SMembers(ctx, strconv.Itoa(int(userId))).Result()
	if err != nil {
		return nil, err
	}
	var res []int64
	for _, id := range ids {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		res = append(res, int64(idInt))
	}
	return res, nil
}

// 从redis中获取用户的关注数
func getFollowingCountFromRedis(ctx context.Context, userId int64) (int64, error) {
	count, err := db.FollowingRedis.SCard(ctx, strconv.Itoa(int(userId))).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 从redis中获取用户的粉丝数
func getFollowersCountFromRedis(ctx context.Context, userId int64) (int64, error) {
	count, err := db.FollowersRedis.SCard(ctx, strconv.Itoa(int(userId))).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}
