package service

import (
	"context"
	"douyin/cmd/favorite/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/video"
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
	res, err := db.GetFavoritesByUserId(req.UserId)
	if err != nil {
		return nil, err
	}
	vids, err := pack.GetVideosByFavorites(s.ctx, res)
	if err != nil {
		return nil, err
	}
	return vids, err
}
