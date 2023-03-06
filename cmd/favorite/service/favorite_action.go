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

// NewFavoriteActionService 创建一个点赞服务
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction 点赞功能实现
func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) error {
	// 根据请求点赞
	userId := req.UserId
	videoId := req.VideoId
	if req.ActionType == 1 {
		return s.Favourite(req, userId, videoId)
	}
	// 根据请求取消点赞
	if req.ActionType == 2 {
		return s.CancelFavorite(req, userId, videoId)
	}
	// 参数不合法
	return errno.ErrBind
}

func (s *FavoriteActionService) CancelFavorite(req *favorite.DouyinFavoriteActionRequest, userId int64, videoId int64) error {
	// 查看缓存里面是否存在有对应的user对象
	uidString := strconv.Itoa(int(userId))
	vidString := strconv.Itoa(int(videoId))
	userCnt, err := db.FavoriteRedis.Exists(s.ctx, uidString).Result()
	if err != nil {
		klog.Fatalf("redis查询用户数据出现错误")
		panic(err)
	}
	if userCnt > 1 {
		klog.Fatalf("用户在redis中不唯一")
		panic(err)
	}
	if userCnt == 1 {
		// favoriteRedis里面存在用户对象
		err := db.FavoriteRedis.HDel(s.ctx, uidString, vidString).Err()
		if err != nil {
			klog.Fatalf("redis删除用户对象字段失败")
			return err
		}
		// 发送消息给MQ
		msg, err := json.Marshal(req)
		if err != nil {
			klog.Fatalf("序列化点赞请求参数失败")
			panic(err)
		}
		FavoriteMq.FavoriteActionSend([]byte(msg))
		return nil
	}
	// favoriteRedis数据库不存在，直接操作数据库
	return db.DeleteFavorite(s.ctx, userId, videoId)
}

func (s *FavoriteActionService) Favourite(req *favorite.DouyinFavoriteActionRequest, userId int64, videoId int64) error {
	// 查看缓存里面是否存在有对应的user对象
	uidString := strconv.Itoa(int(userId))
	vidString := strconv.Itoa(int(videoId))
	userCnt, err := db.FavoriteRedis.Exists(s.ctx, uidString).Result()
	if err != nil {
		klog.Fatalf("redis 查询用户数据出现错误")
		panic(err)
	}
	if userCnt > 1 {
		klog.Fatalf("用户在redis中不唯一")
		panic(err)
	}
	if userCnt == 1 {
		err := s.SetFavouriteToRedis(videoId, userId, uidString, vidString)
		if err != nil {
			klog.Fatalf("set favourite to redis fail")
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
	return db.AddFavorite(s.ctx, userId, videoId)
}

func (s *FavoriteActionService) SetFavouriteToRedis(videoId int64, userId int64, uidString string, vidString string) error {
	// favoriteRedis 里面存在指定的用户对象
	// 尝试从 VideoRedis 里面获取对应的video对象
	var dbVideo = new(db.Video)
	videoRedis := Redis.GetVideoFromRedis(s.ctx, []int64{videoId})
	// 没有对应的video对象
	if videoRedis == nil || videoRedis[0] == nil {
		videoMysql, err := db.GetVideoByIds([]int64{videoId})
		if err != nil {
			klog.Fatalf("数据库获取视频数据失败")
			return err
		}
		// 写入到 VideoRedis 里面
		Redis.PutVideoInRedis(s.ctx, videoMysql[0])
		dbVideo = videoMysql[0]
	} else {
		dbVideo = videoRedis[0]
	}
	// 打包成rpc格式数据
	rpcVideo, err := pack.Video(s.ctx, userId, dbVideo)
	if err != nil {
		klog.Fatalf("视频数据打包成RPC格式出错")
		return err
	}
	// 存储进 FavoriteRedis
	vB, err := json.Marshal(rpcVideo)
	if err != nil {
		klog.Fatalf("RPC格式视频数据序列化出错")
		return err
	}
	err = db.FavoriteRedis.HSet(s.ctx, uidString, vidString, vB).Err()
	if err != nil {
		klog.Fatalf("redis添加用户对象字段失败")
		return err
	}
	return nil
}
