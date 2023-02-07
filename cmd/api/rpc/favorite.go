package rpc

import (
	"context"
	"time"

	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/favorite/favoritesrv"
	"douyin/pkg/errno"
	"github.com/kitex-contrib/registry-nacos/resolver"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var favoriteClient favoritesrv.Client

// favorite客户端初始化
func initFavoriteRpc() {
	hlog.Info("Favorite Client PSM:" + FavoriteRPCPSM)

	c, err := favoritesrv.NewClient(
		FavoriteRPCPSM,
		client.WithResolver(resolver.NewNacosResolver(NacosInit())),
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: FavoriteRPCPSM}),
	)
	if err != nil {
		hlog.Fatal("客户端启动失败")
		panic(err)
	}
	favoriteClient = c
}

// FavoriteAction 传递点赞操作的上下文，同时获取rpc服务端的响应结果
func FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp, err = favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

// FavoriteList 传递获取点赞视频的上下文，同时获取rpc服务端的响应结果
func FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	resp, err = favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
