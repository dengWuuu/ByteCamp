/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-03 22:16:48
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-03 23:25:18
 * @FilePath: /ByteCamp/cmd/relation/service/relation_redis.go
 * @Description: relation微服务对redis的操作封装
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import (
	"context"
	"douyin/dal/db"
	"strconv"
)

//向redis的following set中添加关注
func addRedisFollowList(ctx context.Context, userId, toUserId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FollowingRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		return err
	}
	//如果redis中不存在该用户的关注列表，那么不对redis进行操作
	if cnt != 0 {
		err = db.FollowingRedis.SAdd(ctx, userIdStr, toUserId).Err()
		if err != nil {
			return err
		}
		err = db.FollowingRedis.Expire(ctx, userIdStr, db.ExpireTime).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

//向redis的follower set中添加粉丝
func addRedisFollowerList(ctx context.Context, userId, toUserId int64) error {
	toUserIdStr := strconv.Itoa(int(toUserId))
	cnt, err := db.FollowersRedis.Exists(ctx, toUserIdStr).Result()
	if err != nil {
		return err
	}
	//如果redis中不存在该用户的关注列表，那么不对redis进行操作
	if cnt != 0 {
		err = db.FollowersRedis.SAdd(ctx, toUserIdStr, userId).Err()
		if err != nil {
			return err
		}
		err = db.FollowersRedis.Expire(ctx, toUserIdStr, db.ExpireTime).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

//向redis的friends set中添加朋友
func addRedisFriendsList(ctx context.Context, userId, toUserId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FriendsRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		return err
	}
	//如果redis中不存在该用户的关注列表，那么不对redis进行操作
	if cnt != 0 {
		//TODO:如果要添加朋友关系，首先必须将toUser的following set添加进来

		err = db.FriendsRedis.SAdd(ctx, userIdStr, toUserId).Err()
		if err != nil {
			return err
		}
		err = db.FriendsRedis.Expire(ctx, userIdStr, db.ExpireTime).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

//向redis的following set中移除关注
func remRedisFollowList(ctx context.Context, userId, toUserId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FollowingRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		return err
	}
	//如果redis中不存在该用户的关注列表，那么不对redis进行操作
	if cnt != 0 {
		err = db.FollowingRedis.SRem(ctx, userIdStr, toUserId).Err()
		if err != nil {
			return err
		}
		err = db.FollowingRedis.Expire(ctx, userIdStr, db.ExpireTime).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

//向redis的follower set中移除粉丝
func remRedisFollowerList(ctx context.Context, userId, toUserId int64) error {
	toUserIdStr := strconv.Itoa(int(toUserId))
	cnt, err := db.FollowersRedis.Exists(ctx, toUserIdStr).Result()
	if err != nil {
		return err
	}
	//如果redis中不存在该用户的关注列表，那么不对redis进行操作
	if cnt != 0 {
		err = db.FollowersRedis.SRem(ctx, toUserIdStr, userId).Err()
		if err != nil {
			return err
		}
		err = db.FollowersRedis.Expire(ctx, toUserIdStr, db.ExpireTime).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

//向redis的friends set中添加朋友
func remRedisFriendsList(ctx context.Context, userId, toUserId int64) error {
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FriendsRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		return err
	}
	//如果redis中不存在该用户的关注列表，那么不对redis进行操作
	if cnt != 0 {
		err = db.FriendsRedis.SRem(ctx, userIdStr, toUserId).Err()
		if err != nil {
			return err
		}
		err = db.FriendsRedis.Expire(ctx, userIdStr, db.ExpireTime).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

//从数据库中加载关注列表,并将其存入redis
func loadFollowingListFromDB(ctx context.Context, userId int64) error {
	return nil
}
