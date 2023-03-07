package service

import (
	"context"
	"encoding/json"
	"strconv"

	FavoriteMq "douyin/cmd/favorite/favoriteMq"
	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
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
	//set data to redis
	err := s.CancelFavouriteToRedis(uidString, vidString)
	if err != nil {
		klog.Errorf("cancel favourite to redis fail")
		return err
	}
	// 发送消息给MQ
	msg, err := json.Marshal(req)
	if err != nil {
		klog.Errorf("序列化点赞请求参数失败")
		panic(err)
	}
	FavoriteMq.FavoriteActionSend(msg)
	return nil

}

func (s *FavoriteActionService) Favourite(req *favorite.DouyinFavoriteActionRequest, userId int64, videoId int64) error {
	// 查看缓存里面是否存在有对应的user对象
	uidString := strconv.Itoa(int(userId))
	vidString := strconv.Itoa(int(videoId))

	//set data to redis
	err := s.SetFavouriteToRedis(uidString, vidString)
	if err != nil {
		klog.Errorf("set favourite to redis fail")
		return err
	}

	// 发送消息给MQ
	msg, err := json.Marshal(req)
	if err != nil {
		klog.Errorf("序列化点赞请求参数失败")
		panic(err)
		return err
	}
	FavoriteMq.FavoriteActionSend(msg)
	return nil
}

func (s *FavoriteActionService) SetFavouriteToRedis(uidString string, vidString string) error {
	//increase the favourite nums of the video
	boolCmd := db.FavoriteRedis.SIsMember(s.ctx, vidString, uidString)
	if boolCmd.Val() {
		klog.Errorf("user: %s favourite : %s conflict", uidString, vidString)
		return nil
	}
	add := db.FavoriteRedis.SAdd(s.ctx, vidString, uidString)
	if add.Err() != nil {
		klog.Errorf("user ? set favourite to video ? fail?", vidString, uidString)
		return add.Err()
	}
	return nil
}

func (s *FavoriteActionService) CancelFavouriteToRedis(uidString string, vidString string) error {
	//increase the favourite nums of the video
	boolCmd := db.FavoriteRedis.SIsMember(s.ctx, vidString, uidString)
	if !boolCmd.Val() {
		klog.Errorf("user: %s can't cancel favourite : %s ", uidString, vidString)
		return errno.ErrInvalidValueOfLength
	}
	cancel := db.FavoriteRedis.SRem(s.ctx, vidString, uidString)
	if cancel.Err() != nil {
		klog.Errorf("user ? cancel favourite to video ? fail", vidString, uidString)
		return cancel.Err()
	}
	return nil
}
