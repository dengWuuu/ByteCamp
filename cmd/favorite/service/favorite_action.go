package service

import (
	"context"

	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
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
		return db.AddFavorite(s.ctx, user_id, video_id)
	}
	// 根据请求取消点赞
	if req.ActionType == 2 {
		return db.DeleteFavorite(s.ctx, user_id, video_id)
	}
	// 参数不合法
	return errno.ErrBind
}
