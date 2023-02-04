package main

import (
	"douyin/dal"
	"douyin/kitex_gen/video/videosrv"
	"douyin/pkg/nacos"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
)

// Init Relation RPC Server 端配置初始化
func Init() {
	dal.Init()
}
func main() {
	Init()

	PSM := "bytecamp.douyin.video"
	Address := "127.0.0.1"
	Port := 8084
	//Port, err := nacos.GetFreePort()
	//if err != nil{
	//	panic(err)
	//}
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", Address, Port))

	//nacos
	r := registry.NewNacosRegistry(nacos.InitNacos())
	svr := videosrv.NewServer(new(VideoSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: PSM}))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
