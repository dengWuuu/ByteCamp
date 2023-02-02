package rpc

import (
	"context"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/favorite/favoritesrv"
	"douyin/pkg/errno"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/spf13/viper"
)

var favoriteClient favoritesrv.Client

// favorite客户端初始化
func initFavoriteRpc() {
	//读取配置
	path, err1 := os.Getwd()
	if err1 != nil {
		panic(err1)
	}
	viper.SetConfigName("favoriteService")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config")
	errV := viper.ReadInConfig()
	if errV != nil {
		hlog.Fatal("启动rpc favorite服务器时读取配置文件失败")
		return
	}
	commentSrvPath := viper.GetString("Server.Address") + ":" + viper.GetString("Server.Port")
	hlog.Info("favorite客户端对应的服务端地址" + commentSrvPath)
	c, err := favoritesrv.NewClient(
		viper.GetString("Server.Name"),
		client.WithHostPorts(commentSrvPath),
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: viper.GetString("Server.Name")}),
	)
	if err != nil {
		hlog.Fatal("客户端启动失败")
		panic(err)
	}
	favoriteClient = c
}

// 传递点赞操作的上下文，同时获取rpc服务端的响应结果
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

// 传递获取点赞视频的上下文，同时获取rpc服务端的响应结果
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
