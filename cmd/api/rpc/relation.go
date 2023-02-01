/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 02:20:30
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 13:13:48
 * @FilePath: /ByteCamp/cmd/api/rpc/relation.go
 * @Description: 用于初始化relation微服务的client,并且通过relation微服务的client调用relation微服务的方法从而实现api中http接口
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package rpc

import (
	"context"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationsrv"
	"douyin/pkg/errno"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/spf13/viper"
)

var relationClient relationsrv.Client

// init 初始化relation rpc 客户端
func initRelationRpc() {
	//读取配置
	path, err1 := os.Getwd()
	if err1 != nil {
		panic(err1)
	}
	viper.SetConfigName("relationService")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config")
	errV := viper.ReadInConfig()
	if errV != nil {
		hlog.Fatal("启动rpc客户端时读取配置文件失败")
		return
	}
	relationSrvPath := viper.GetString("Server.Address") + ":" + viper.GetString("Server.Port")
	hlog.Info("relation客户端对应的服务端地址" + relationSrvPath)
	c, err := relationsrv.NewClient(
		viper.GetString("Server.Name"),
		client.WithHostPorts(relationSrvPath),
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
	relationClient = c
}

func RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	//1、调用rpc接口完成操作,注意需要判断RPC调用是否成功
	resp, err = relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	//2、检查resp是否合法
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
