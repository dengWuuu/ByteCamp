/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-29 21:58:00
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-08 14:27:20
 * @FilePath: /ByteCamp/cmd/relation/main.go
 * @Description: relation rpc server 启动入口
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"fmt"
	"log"
	"net"

	"douyin/cmd/relation/relationMq"
	"douyin/dal"
	relation "douyin/kitex_gen/relation/relationsrv"
	"douyin/pkg/nacos"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"github.com/kitex-contrib/registry-nacos/registry"
)

// Init Relation RPC Server 端配置初始化
func Init() {
	dal.Init()
	relationMq.InitRelationMq() // 初始化mq
}

func main() {
	Init()

	PSM := "bytecamp.douyin.relation"
	Address := "127.0.0.1"
	Port := 8082
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", Address, Port)) // nacos

	// nacos
	r := registry.NewNacosRegistry(nacos.InitNacos())
	svr := relation.NewServer(
		new(RelationSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 100000000000000, MaxQPS: 1000000000}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: PSM}))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
