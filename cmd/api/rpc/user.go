package rpc

import (
	"context"
	"time"

	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/usersrv"
	"douyin/pkg/errno"
	"douyin/pkg/jaeger"
	"douyin/pkg/prometheus"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/registry-nacos/resolver"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
)

var userClient usersrv.Client

// init 初始化用户 rpc 客户端
func initUserRpc() {
	hlog.Info("User Client PSM:" + UserRPCPSM)

	tracerSuite, _ := jaeger.InitJaegerClient("user-client")

	c, err := usersrv.NewClient(
		UserRPCPSM,
		client.WithTracer(prometheus.KitexClientTracer),
		client.WithResolver(resolver.NewNacosResolver(NacosInit())),
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: UserRPCPSM}),
		client.WithSuite(tracerSuite),
	)
	if err != nil {
		hlog.Fatal("客户端启动失败")
		panic(err)
	}
	userClient = c
}

// Register 注册方法，传递注册上下文，并且获取prc响应
func Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp, err = userClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp, err = userClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp, err = userClient.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
