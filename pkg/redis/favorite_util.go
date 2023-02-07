package redis

import (
	"context"
	"douyin/dal/db"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
)

// 判断用户 user_id 是否点赞了视频 video_id
func IsFavorite(ctx context.Context, user_id int64, video_id int64) bool {
	// 首先判断是否能够找到favoriteRedis的缓存
	uid_string := strconv.Itoa(int(user_id))
	vid_string := strconv.Itoa(int(video_id))
	user_cnt, err := db.FavoriteRedis.Exists(ctx, uid_string).Result()
	if err != nil {
		klog.Error("redis数据库查询异常")
		panic(err)
	}
	if user_cnt > 1 {
		klog.Error("指定的用户不唯一")
		panic(err)
	}
	if user_cnt == 1 {
		// favoriteRedis缓存里面有
		result, err := db.FavoriteRedis.HExists(ctx, uid_string, vid_string).Result()
		if err != nil {
			klog.Error("redis数据库执行异常")
			panic(err)
		}
		return result
	}
	// 缓存没有，只能查询数据库
	fs, err := db.GetFavoritesByUserId(user_id)
	if err != nil {
		klog.Error("mysql数据查询数据异常")
		panic(err)
	}
	for _, f := range fs {
		if int64(f.VideoId) == video_id {
			return true
		}
	}
	return false
}
