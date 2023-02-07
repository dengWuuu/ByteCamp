package pack

import (
	"context"

	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/user"
	Redis "douyin/pkg/redis"
)

// 包装db层数据结构到rpc通信数据结构
func Comment(ctx context.Context, m *db.Comment) (*comment.Comment, error) {
	// 检查非空
	if m == nil {
		return &comment.Comment{Content: "no content"}, nil
	}
	// 获取对应的评论用户数据
	var u *db.User
	redis_user := Redis.GetUsersFromRedis(ctx, []uint{uint(m.UserId)})
	if redis_user == nil {
		db_u, err := db.GetUserById(int64(m.UserId))
		if err != nil {
			return nil, err
		}
		u = db_u
	} else {
		u = redis_user[0]
	}
	// 转换用户字段数据的类型
	followCount := int64(u.FollowingCount)
	followerCount := int64(u.FollowerCount)
	us := &user.User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      false,
	}
	// 转换评论类型
	// * 记得将时间格式化
	return &comment.Comment{
		Id:         int64(m.ID),
		User:       us,
		Content:    m.Content,
		CreateDate: m.CreatTime.Format("2006-01-02 15:04:05"),
	}, nil
}

func Comments(ctx context.Context, ms []*db.Comment) ([]*comment.Comment, error) {
	comments := make([]*comment.Comment, 0)
	for _, m := range ms {
		if n, err := Comment(ctx, m); n != nil && err == nil {
			comments = append(comments, n)
		}
	}
	return comments, nil
}
