package main

import (
	"fmt"
	"log"
	"net"

	"douyin/cmd/video/config"
	"douyin/cmd/video/dal"
	"douyin/kitex_gen/video/videosrv"
	"douyin/pkg/nacos"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"

	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
)

// Init Relation RPC Server 端配置初始化
func Init() {
	dal.Init()
}

func main() {
	Init()

	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.Address, config.Port))

	svr := videosrv.NewServer(
		new(VideoSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry.NewNacosRegistry(nacos.InitNacos())),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.PSM}))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
