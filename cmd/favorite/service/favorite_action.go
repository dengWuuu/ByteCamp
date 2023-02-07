package service

import (
	"context"
	"encoding/json"
	"strconv"

	FavoriteMq "douyin/cmd/favorite/favoriteMq"
	"douyin/cmd/favorite/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	Redis "douyin/pkg/redis"

	"github.com/cloudwego/kitex/pkg/klog"
)

type FavoriteActionService struct {
	ctx context.Context
}

// 创建一个点赞服务
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// 点赞功能实现
func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) error {
	// 根据请求点赞
	user_id := req.UserId
	video_id := req.VideoId
	if req.ActionType == 1 {
		// 查看缓存里面是否存在有对应的user对象
		uid_string := strconv.Itoa(int(user_id))
		vid_string := strconv.Itoa(int(video_id))
		user_cnt, err := db.FavoriteRedis.Exists(s.ctx, uid_string).Result()
		if err != nil {
			klog.Fatalf("redis 查询用户数据出现错误")
			panic(err)
		}
		if user_cnt > 1 {
			klog.Fatalf("用户在redis中不唯一")
			panic(err)
		}
		if user_cnt == 1 {
			// favoriteRedis 里面存在指定的用户对象
			// 尝试从 VideoRedis 里面获取对应的video对象
			var dbVideo = new(db.Video)
			video_redis := Redis.GetVideoFromRedis(s.ctx, []int64{video_id})
			// 没有对应的video对象
			if video_redis == nil || video_redis[0] == nil {
				video_mysql, err := db.GetVideoByIds([]int64{video_id})
				if err != nil {
					klog.Fatalf("数据库获取视频数据失败")
					return err
				}
				// 写入到 VideoRedis 里面
				Redis.PutVideoInRedis(s.ctx, video_mysql[0])
				dbVideo = video_mysql[0]
			} else {
				dbVideo = video_redis[0]
			}
			// 打包成rpc格式数据
			rpcVideo, err := pack.Video(s.ctx, dbVideo)
			if err != nil {
				klog.Fatalf("视频数据打包成RPC格式出错")
				return err
			}
			// 存储进 FavoriteRedis
			v_b, err := json.Marshal(rpcVideo)
			if err != nil {
				klog.Fatalf("RPC格式视频数据序列化出错")
				return err
			}
			err = db.FavoriteRedis.HSet(s.ctx, uid_string, vid_string, v_b).Err()
			if err != nil {
				klog.Fatalf("redis添加用户对象字段失败")
				return err
			}
			// * 发送消息给MQ
			msg, err := json.Marshal(req)
			if err != nil {
				klog.Fatalf("序列化点赞请求参数失败")
				panic(err)
			}
			FavoriteMq.FavoriteActionSend([]byte(msg))
			return nil
		}
		// favoriteRedis里面没有对应的数据
		// 直接修改数据库
		return db.AddFavorite(s.ctx, user_id, video_id)
	}
	// 根据请求取消点赞
	if req.ActionType == 2 {
		// 查看缓存里面是否存在有对应的user对象
		uid_string := strconv.Itoa(int(user_id))
		vid_string := strconv.Itoa(int(video_id))
		user_cnt, err := db.FavoriteRedis.Exists(s.ctx, uid_string).Result()
		if err != nil {
			klog.Fatalf("redis查询用户数据出现错误")
			panic(err)
		}
		if user_cnt > 1 {
			klog.Fatalf("用户在redis中不唯一")
			panic(err)
		}
		if user_cnt == 1 {
			// favoriteRedis里面存在用户对象
			err := db.FavoriteRedis.HDel(s.ctx, uid_string, vid_string).Err()
			if err != nil {
				klog.Fatalf("redis删除用户对象字段失败")
				return err
			}
			// * 发送消息给MQ
			msg, err := json.Marshal(req)
			if err != nil {
				klog.Fatalf("序列化点赞请求参数失败")
				panic(err)
			}
			FavoriteMq.FavoriteActionSend([]byte(msg))
			return nil
		}
		// favoriteRedis数据库不存在，直接操作数据库
		return db.DeleteFavorite(s.ctx, user_id, video_id)
	}
	// 参数不合法
	return errno.ErrBind
}
