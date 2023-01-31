package pack

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/user"
)

// 包装db层数据结构到rpc通信数据结构
func Comment(ctx context.Context, m *db.Comment) (*comment.Comment, error) {
	// 检查非空
	if m == nil {
		return &comment.Comment{Content: "no content"}, nil
	}
	// 获取对应的用户数据
	u, err := db.GetUserById(m.UserId)
	if err != nil || u == nil {
		return &comment.Comment{Content: "no user"}, nil
	}
	// 转换用户的类型
	followCount := int64(u.FollowingCount)
	followerCount := int64(u.FollowerCount)
	user := &user.User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      false,
	}
	// 转换评论类型
	return &comment.Comment{
		Id:         int64(m.ID),
		User:       user,
		Content:    m.Content,
		CreateDate: m.CreatTime.String(),
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
