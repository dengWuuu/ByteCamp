package service

import (
	"context"
	"encoding/json"
	"strconv"

	"douyin/cmd/favorite/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/video"
	Redis "douyin/pkg/redis"

	"github.com/cloudwego/kitex/pkg/klog"
)

type FavoriteListService struct {
	ctx context.Context
}

// 创建一个获取点赞视频服务
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

// 实现服务具体功能
func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*video.Video, error) {
	// 判断缓存里面是否有对象
	user_id := req.UserId
	user_string := strconv.Itoa(int(user_id))
	user_cnt, err := db.FavoriteRedis.Exists(s.ctx, user_string).Result()
	if err != nil {
		klog.Fatalf("redis查找对象出错")
		panic(err)
	}
	if user_cnt > 1 {
		klog.Fatalf("video对象不唯一")
		panic(err)
	}
	if user_cnt == 1 {
		// 存在user对象
		video_data, err := db.FavoriteRedis.HGetAll(s.ctx, user_string).Result()
		if err != nil {
			panic(err)
		}
		var videos = make([]*video.Video, 0)
		for _, video_redis := range video_data {
			var vid = new(video.Video)
			err := json.Unmarshal([]byte(video_redis), vid)
			if err != nil {
				klog.Fatalf("redis 获取视频数据序列化失败")
				panic(err)
			}
			videos = append(videos, vid)
		}
		return videos, nil
	}
	// redis缓存没有对应的数据，需要从数据库读取
	// ^ 分为两个步骤：获取video信息，再打包成rpc的数据结构信息
	resp, err := db.GetFavoritesByUserId(user_id)
	if err != nil {
		klog.Fatalf("mysql 获取用户喜欢视频列表失败")
		panic(err)
	}
	// 用户喜欢的视频列表ID
	var vids = make([]int64, len(resp))
	for index, f := range resp {
		vids[index] = int64(f.VideoId)
	}
	video_redis := Redis.GetVideoFromRedis(s.ctx, vids)
	if video_redis == nil {
		// redis里面没有对应的video信息
		// 需要从数据库中调取
		video_mysql, err := db.GetVideoByIds(vids)
		if err != nil {
			klog.Fatalf("mysql 获取视频信息出错")
			panic(err)
		}
		video_data, err := pack.Videos(s.ctx, video_mysql)
		if err != nil {
			klog.Fatalf("将数据库打包成rpc数据格式出错")
			panic(err)
		}
		// 把视频信息加入到video redis里面
		for _, vid_data := range video_mysql {
			Redis.PutVideoInRedis(s.ctx, vid_data)
		}
		// 把用户的结果放到favorite redis里面
		err = AddVideoListInRedis(s.ctx, user_id, video_data)
		if err != nil {
			panic(err)
		}
		return video_data, nil
	}
	// 查找其中的不存在的视频信息
	var vids_nil = make([]int64, 0)
	for index, v := range video_redis {
		if v == nil {
			vids_nil = append(vids_nil, vids[index])
		}
	}
	// 从数据库中获取不存在的视频信息并且存储到video redis里面
	if len(vids_nil) > 0 {
		vid_mysql, err := db.GetVideoByIds(vids_nil)
		if err != nil {
			klog.Fatalf("mysql 获取视频信息出错")
			panic(err)
		}
		// 把视频信息加入到video redis里面
		for _, vid_data := range vid_mysql {
			Redis.PutVideoInRedis(s.ctx, vid_data)
		}
		// 重新查一遍
		video_redis = Redis.GetVideoFromRedis(s.ctx, vids)
	}
	video_data, err := pack.Videos(s.ctx, video_redis)
	if err != nil {
		klog.Fatalf("将数据库打包成rpc数据格式出错")
		panic(err)
	}
	// 把用户的结果放到favorite redis里面
	err = AddVideoListInRedis(s.ctx, user_id, video_data)
	if err != nil {
		panic(err)
	}
	return video_data, nil
}
