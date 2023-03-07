package main

import (
	"context"
	"douyin/cmd/favorite/pack"
	"douyin/cmd/favorite/service"
	favorite "douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

// FavoriteSrvImpl implements the last service interface defined in the IDL.
type FavoriteSrvImpl struct{}

// FavoriteAction implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	favoriteSrv := service.NewFavoriteActionService(ctx)
	// 检查参数
	if req.UserId <= 0 || req.VideoId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		return pack.BuildFavoriteActionResp(errno.ErrBind), nil
	}
	err = favoriteSrv.FavoriteAction(req)
	return pack.BuildFavoriteActionResp(err), nil
}

// FavoriteList implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	favoriteSrv := service.NewFavoriteListService(ctx)
	// 检查参数
	if req.UserId <= 0 {
		return pack.BuildFavoriteListResp(errno.ErrBind), nil
	}
	vids, err := favoriteSrv.FavoriteList(req)
	if err != nil {
		return nil, err
	}
	packRes := pack.BuildFavoriteListResp(err)
	packRes.VideoList = vids
	return packRes, err
}
