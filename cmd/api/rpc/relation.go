/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 02:20:30
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 18:32:37
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
	"github.com/kitex-contrib/registry-nacos/resolver"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var relationClient relationsrv.Client

// init 初始化relation rpc 客户端
func initRelationRpc() {
	hlog.Info("Relation Client PSM:" + RelationRPCPSM)

	c, err := relationsrv.NewClient(
		RelationRPCPSM,
		client.WithResolver(resolver.NewNacosResolver(NacosInit())),
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: RelationRPCPSM}),
	)
	if err != nil {
		hlog.Fatal("客户端启动失败")
		panic(err)
	}
	relationClient = c
}

// RelationAction 用户关注或取消关注
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

// FollowList 用户关注列表
func FollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	resp, err = relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

// FollowerList 用户粉丝列表
func FollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	resp, err = relationClient.RelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func FriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	resp, err = relationClient.RelationFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
