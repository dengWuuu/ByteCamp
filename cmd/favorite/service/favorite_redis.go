package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/video"
	"encoding/json"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
)

// AddVideoListInRedis key: user_id, value: map[video_id]{video.Video}
func AddVideoListInRedis(ctx context.Context, uid int64, vs []*video.Video) error {
	videoMap := make(map[string]interface{})
	for _, v := range vs {
		if v != nil {
			vBinary, err := json.Marshal(v)
			if err != nil {
				klog.Fatalf("视频数据序列化失败")
				return err
			}
			vString := strconv.Itoa(int(v.Id))
			videoMap[vString] = vBinary
		}
	}
	// ! 一定要判断是否为空
	if len(videoMap) == 0 {
		return nil
	}
	uString := strconv.Itoa(int(uid))
	err := db.FavoriteRedis.HMSet(ctx, uString, videoMap).Err()
	if err != nil {
		klog.Fatalf("视频数据写入redis失败")
		return err
	}
	// 设置过期时间
	db.FavoriteRedis.Expire(ctx, uString, db.ExpireTime)
	return nil
}
