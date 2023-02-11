package main

import (
	"fmt"
	"log"
	"net"

	"douyin/dal"
	user "douyin/kitex_gen/user/usersrv"
	"douyin/pkg/jaeger"
	"douyin/pkg/nacos"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"github.com/kitex-contrib/registry-nacos/registry"
)

// Init User RPC Server 端配置初始化
func Init() {
	dal.Init()
}

func main() {
	Init()

	PSM := "bytecamp.douyin.user"
	Address := "127.0.0.1"
	Port := 8081
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", Address, Port)) // nacos
	// nacos
	r := registry.NewNacosRegistry(nacos.InitNacos())

	//jaeger
	tracerSuite, closer := jaeger.InitJaegerServer("user-server")
	defer closer.Close()

	svr := user.NewServer(
		new(UserSrvImpl),
		server.WithTracer(prometheus.NewServerTracer(":9094", "/metrics")),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: PSM}),
		server.WithSuite(tracerSuite),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
