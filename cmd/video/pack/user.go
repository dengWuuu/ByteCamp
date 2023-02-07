package pack

import (
	"context"

	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/user"
)

// User 包装数据库的数据成为rpc中用的数据
func User(ctx context.Context, u *db.User) (*user.User, error) {
	if u == nil {
		return &user.User{
			Name: "无此用户",
		}, nil
	}

	followCount := int64(u.FollowingCount)
	followerCount := int64(u.FollowerCount)

	isFollow := false
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
	}, nil
}
