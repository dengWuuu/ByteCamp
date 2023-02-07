package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/video"
	"encoding/json"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
)

// key: user_id, value: map[video_id]{video.Video}
func AddVideoListInRedis(ctx context.Context, uid int64, vs []*video.Video) error {
	video_map := make(map[string]interface{})
	for _, v := range vs {
		if v != nil {
			v_binary, err := json.Marshal(v)
			if err != nil {
				klog.Fatalf("视频数据序列化失败")
				return err
			}
			v_string := strconv.Itoa(int(v.Id))
			video_map[v_string] = v_binary
		}
	}
	// ! 一定要判断是否为空
	if len(video_map) == 0 {
		return nil
	}
	u_string := strconv.Itoa(int(uid))
	err := db.FavoriteRedis.HMSet(ctx, u_string, video_map).Err()
	if err != nil {
		klog.Fatalf("视频数据写入redis失败")
		return err
	}
	// 设置过期时间
	db.FavoriteRedis.Expire(ctx, u_string, db.ExpireTime)
	return nil
}
