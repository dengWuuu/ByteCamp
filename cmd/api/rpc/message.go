package rpc

import (
	"context"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/message/messagesrv"
	"douyin/pkg/errno"
	"douyin/pkg/jaeger"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"time"
)

var messageClient messagesrv.Client

func initMessageRpc() {
	hlog.Info("Message Client PSM:" + MessageRPCPSM)

	//jaeger
	tracerSuite, _ := jaeger.InitJaegerClient("message-client")

	c, err := messagesrv.NewClient(
		MessageRPCPSM,
		client.WithResolver(resolver.NewNacosResolver(NacosInit())),
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: MessageRPCPSM}),
		client.WithSuite(tracerSuite),
	)
	if err != nil {
		hlog.Fatal("客户端启动失败")
		panic(err)
	}
	messageClient = c
}

func MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	// 1、调用rpc接口完成操作,注意需要判断RPC调用是否成功
	resp, err = messageClient.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	// 2、检查resp是否合法
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
