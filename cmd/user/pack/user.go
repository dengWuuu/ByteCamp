package pack

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/redis"
)

// User 包装数据库的数据成为rpc中用的数据
func User(ctx context.Context, u *db.User, fromId int64) (*user.User, error) {
	if u == nil {
		return &user.User{
			Name: "无此用户",
		}, nil
	}

	followCount := int64(u.FollowingCount)
	followerCount := int64(u.FollowerCount)
	isFollow, err := redis.IsFollowing(ctx, fromId, int64(u.ID))
	if err != nil {
		return nil, err
	}

	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
	}, nil
}

// Users pack list of userHandler info
func Users(ctx context.Context, us []*db.User, fromId int64) ([]*user.User, error) {
	users := make([]*user.User, 0)
	for _, u := range us {
		user2, err := User(ctx, u, fromId)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}
