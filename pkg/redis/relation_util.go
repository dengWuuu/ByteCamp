package redis

import (
	"context"
	"strconv"

	"douyin/dal/db"
)

// 判断一个用户是否关注了另一个用户
func IsFollowing(ctx context.Context, userId, otherId int64) (bool, error) {
	// 1.查看redis中是否有缓存
	userIdStr := strconv.Itoa(int(userId))
	cnt, err := db.FollowingRedis.Exists(ctx, userIdStr).Result()
	if err != nil {
		return false, err
	}
	if cnt == 0 {
		// loadFollowingListToRedis(ctx, userId)
		// copy of loadFollowingListToRedis,防止循环引用
		err := db.FollowingRedis.SAdd(ctx, strconv.Itoa(int(userId)), -1).Err()
		if err != nil {
			return false, err
		}
		ids, err := db.GetFollowingByUserId(int(userId))
		if err != nil {
			return false, err
		}
		var idsStr []string
		for _, id := range ids {
			idsStr = append(idsStr, strconv.Itoa(int(id)))
		}
		err = db.FollowingRedis.SAdd(ctx, strconv.Itoa(int(userId)), idsStr).Err()
		if err != nil {
			return false, err
		}
		err = db.FollowingRedis.Expire(ctx, strconv.Itoa(int(userId)), db.ExpireTime).Err()
		if err != nil {
			return false, err
		}
	}
	// 2.判断redis缓存中是否有otherId
	otherIdStr := strconv.Itoa(int(otherId))
	exists, err := db.FollowingRedis.SIsMember(ctx, userIdStr, otherIdStr).Result()
	if err != nil {
		return false, err
	}
	return exists, nil
}
