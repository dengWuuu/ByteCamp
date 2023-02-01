package rpc

import (
	"context"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/comment/commentsrv"
	"douyin/pkg/errno"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/spf13/viper"
)

var commentClient commentsrv.Client

// comment客户端初始化
func initCommentRpc() {
	//读取配置
	path, err1 := os.Getwd()
	if err1 != nil {
		panic(err1)
	}
	viper.SetConfigName("commentService")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config")
	errV := viper.ReadInConfig()
	if errV != nil {
		hlog.Fatal("启动rpc用户服务器时读取配置文件失败")
		return
	}
	commentSrvPath := viper.GetString("Server.Address") + ":" + viper.GetString("Server.Port")
	hlog.Info("comment客户端对应的服务端地址" + commentSrvPath)
	c, err := commentsrv.NewClient(
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
	commentClient = c
}

// 传递评论操作的上下文，并且获取RPC服务端的响应
func CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp, err = commentClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

// 传递评论获取的上下文，并且获取RPC服务端的响应
func CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp, err = commentClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
