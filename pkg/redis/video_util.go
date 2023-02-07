package redis

import (
	"context"
	"douyin/dal/db"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
)

const video_prefix = "video:Id:"

// 从redis里面获取视频信息
func GetVideoFromRedis(ctx context.Context, videoIds []int64) []*db.Video {
	var videoList = make([]*db.Video, len(videoIds))
	video_pip := db.VideoRedis.Pipeline()

	// * 使用管道命令一次性传输多条redis命令，减少IO
	for i := 0; i < len(videoIds); i++ {
		video_pip.Get(ctx, video_prefix+strconv.Itoa(int(videoIds[i])))
	}
	res, err := video_pip.Exec(ctx)
	if err != nil {
		// 没有找到，返回nil
		klog.Error("Redis 执行命令失败")
		return nil
	}
	for index, cmdRes := range res {
		cmd, ok := cmdRes.(*redis.StringCmd)
		if !ok {
			klog.Fatalf("Redis数据转换失败")
			return nil
		}
		bytes, err := cmd.Bytes()
		if err != nil {
			klog.Fatalf("redis获取视频信息失败,获取字节码数组失败")
			videoList[index] = nil
			continue
		}
		vid := new(db.Video)
		err = json.Unmarshal(bytes, vid)
		if err != nil {
			klog.Fatalf("redis中获取视频信息后反序列化失败")
			return nil
		}
		videoList[index] = vid
	}
	return videoList
}

// 把视频信息放到redis里面
func PutVideoInRedis(ctx context.Context, video *db.Video) {
	video_binary, err := json.Marshal(video)
	if err != nil {
		klog.Fatalf("添加到Redis过程序列化视频数据失败")
	}
	vid_key := video_prefix + strconv.Itoa(int(video.ID))
	res, err := db.VideoRedis.Set(ctx, vid_key, video_binary, time.Hour*48).Result()
	if err != nil {
		klog.Error("Redis放进视频" + res + "失败")
	} else {
		klog.Info("Redis放进视频" + res + "成功")
	}
}
