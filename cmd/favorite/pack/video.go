package pack

import (
	"context"

	"douyin/dal/db"
	"douyin/kitex_gen/video"
	Redis "douyin/pkg/redis"
)

// Video 封装db层数据结构为rpc使用的数据结构
func Video(ctx context.Context, userId int64, m *db.Video) (*video.Video, error) {
	// 检查非空
	if m == nil {
		return nil, nil
	}
	var u *db.User
	redisUser := Redis.GetUsersFromRedis(ctx, []uint{uint(m.AuthorId)})
	if redisUser == nil {
		dbU, err := db.GetUserById(m.AuthorId)
		if err != nil {
			return nil, err
		}
		u = dbU
	} else {
		u = redisUser[0]
	}
	// 打包用户的数据
	// * 重点在于检查是否关注
	author, err := User(ctx, userId, u)
	if err != nil {
		return nil, err
	}
	// * 检查是否已经点赞
	isFavorite := Redis.IsFavorite(ctx, userId, int64(m.ID))

	return &video.Video{
		Id:            int64(m.ID),
		Author:        author,
		PlayUrl:       m.PlayUrl,
		CoverUrl:      m.CoverUrl,
		FavoriteCount: m.FavoriteCount,
		CommentCount:  m.CommentCount,
		Title:         m.Title,
		IsFavorite:    isFavorite,
	}, nil
}

func Videos(ctx context.Context, userId int64, ms []*db.Video) ([]*video.Video, error) {
	videos := make([]*video.Video, 0)
	for _, m := range ms {
		if n, err := Video(ctx, userId, m); err == nil && n != nil {
			videos = append(videos, n)
		}
	}
	return videos, nil
}
