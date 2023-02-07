package pack

import (
	"context"

	"douyin/dal/db"
	"douyin/kitex_gen/video"
)

// 包装db层数据结构到rpc通信使用的数据结构
func GetVideosByFavorites(ctx context.Context, m []*db.Favorite) ([]*video.Video, error) {
	videos := make([]int64, 0)
	for _, f := range m {
		videos = append(videos, int64(f.VideoId))
	}
	dbvids, err := db.GetVideoByIds(videos)
	if err != nil {
		return nil, err
	}
	vids, err := Videos(ctx, dbvids)
	if err != nil {
		return nil, err
	}
	return vids, nil
}
