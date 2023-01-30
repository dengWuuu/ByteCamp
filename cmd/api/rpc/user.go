package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/usersrv"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/spf13/viper"
	"os"
	"time"
)

var userClient usersrv.Client

// init 初始化用户 rpc 客户端
func initUserRpc() {
	//读取配置
	path, err1 := os.Getwd()
	if err1 != nil {
		panic(err1)
	}
	viper.SetConfigName("userService")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "\\config")
	errV := viper.ReadInConfig()
	if errV != nil {
		hlog.Fatal("启动rpc用户服务器时读取配置文件失败")
		return
	}
	userSrvPath := viper.GetString("Server.Address") + ":" + viper.GetString("Server.Port")
	hlog.Info("user客户端对应的服务端地址" + userSrvPath)
	c, err := usersrv.NewClient(
		viper.GetString("Server.Name"),
		client.WithHostPorts(userSrvPath),
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
	userClient = c
}

// Register 注册方法，传递注册上下文，并且获取prc响应
func Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp, err = userClient.Register(ctx, req)
	return resp, err
}
