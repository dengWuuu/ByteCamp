package pack

import (
	"context"
	"douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/video"
	"errors"

	"gorm.io/gorm"
)

// Video 封装db层数据结构为rpc使用的数据结构
func Video(ctx context.Context, m *db.Video) (*video.Video, error) {
	// 检查非空
	if m == nil {
		return nil, nil
	}
	user, err := db.GetUserById(int64(m.AuthorId))
	// 检查用户
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	author, err := pack.User(ctx, user)
	if err != nil {
		return nil, err
	}
	favorite_count := int64(m.FavoriteCount)
	comment_count := int64(m.CommentCount)
	return &video.Video{
		Id:            int64(m.ID),
		Author:        author,
		PlayUrl:       m.PlayUrl,
		CoverUrl:      m.CoverUrl,
		FavoriteCount: favorite_count,
		CommentCount:  comment_count,
		Title:         m.Title,
	}, nil
}

func Videos(ctx context.Context, ms []*db.Video) ([]*video.Video, error) {
	videos := make([]*video.Video, 0)
	for _, m := range ms {
		if n, err := Video(ctx, m); err == nil && n != nil {
			videos = append(videos, n)
		}
	}
	return videos, nil
}
