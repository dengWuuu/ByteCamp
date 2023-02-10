package pack

import (
	"context"
	"douyin/dal/db"
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

	// true->fromID已关注u.ID，false-fromID未关注u.ID
	isFollow := false
	//relation, err := db.GetRelation(ctx, fromID, int64(u.ID))
	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	return nil, err
	//}
	//
	//if relation != nil {
	//	isFollow = true
	//}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
	}, nil
}

// Users pack list of userHandler info
func Users(ctx context.Context, us []*db.User) ([]*user.User, error) {
	users := make([]*user.User, 0)
	for _, u := range us {
		user2, err := User(ctx, u)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}
