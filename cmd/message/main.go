package main

import (
	"douyin/pkg/jaeger"
	"fmt"
	"log"
	"net"

	"douyin/cmd/message/config"
	"douyin/dal"
	message "douyin/kitex_gen/message/messagesrv"
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
}

func main() {
	Init()
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.Address, config.Port))

	// jaeger
	tracerSuite, closer := jaeger.InitJaegerServer("message-server")
	defer closer.Close()

	svr := message.NewServer(new(MessageSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry.NewNacosRegistry(nacos.InitNacos())),
		server.WithLimit(&limit.Option{MaxConnections: 100000000, MaxQPS: 1000000000}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.PSM}),
		server.WithSuite(tracerSuite),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
