/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-29 21:58:00
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 14:02:20
 * @FilePath: /ByteCamp/cmd/relation/main.go
 * @Description: relation rpc server 启动入口
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"douyin/dal"
	relation "douyin/kitex_gen/relation/relationsrv"
	"douyin/pkg/nacos"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"github.com/spf13/viper"
)

// Init Relation RPC Server 端配置初始化
func Init() {
	dal.Init()
}
func main() {
	Init()
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
		hlog.Fatal("启动rpc relation 服务器时读取配置文件失败")
		return
	}

	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", viper.GetString("Server.Address")+":"+viper.GetString("Server.Port"))

	//nacos
	r := registry.NewNacosRegistry(nacos.InitNacos())
	svr := relation.NewServer(new(RelationSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: viper.GetString("Server.Name")}))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
