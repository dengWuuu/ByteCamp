package pack

import (
	"context"

	"douyin/dal/db"
	"douyin/kitex_gen/user"
	Redis "douyin/pkg/redis"

	"github.com/cloudwego/kitex/pkg/klog"
)

// User 包装数据库的数据成为rpc中用的数据
func User(ctx context.Context, user_id int64, u *db.User) (*user.User, error) {
	if u == nil {
		return &user.User{
			Name: "无此用户",
		}, nil
	}

	followCount := int64(u.FollowingCount)
	followerCount := int64(u.FollowerCount)

	// true->fromID已关注u.ID，false-fromID未关注u.ID
	isFollow, err := Redis.IsFollowing(ctx, user_id, int64(u.ID))
	if err != nil {
		klog.Error("从redis获取用户关注关系失败")
		panic(err)
	}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
	}, nil
}
