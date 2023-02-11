package main

import (
	"fmt"
	"log"
	"net"

	FavoriteMq "douyin/cmd/favorite/favoriteMq"
	"douyin/dal"
	"douyin/kitex_gen/favorite/favoritesrv"
	"douyin/pkg/jaeger"
	"douyin/pkg/nacos"
	"douyin/pkg/rabbitmq"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"github.com/kitex-contrib/registry-nacos/registry"
)

// Init Relation RPC Server 端配置初始化
func Init() {
	dal.Init()
	FavoriteMq.InitFavoriteMq()
	go FavoriteMq.FavoriteConsumer()
}

func main() {
	Init()
	PSM := "bytecamp.douyin.favorite"
	Address := "127.0.0.1"
	Port := 8085
	//Port, err := nacos.GetFreePort()
	//if err != nil{
	//	panic(err)
	//}
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", Address, Port)) // nacos
	r := registry.NewNacosRegistry(nacos.InitNacos())

	//jaeger
	tracerSuite, closer := jaeger.InitJaegerServer("favorite-server")
	defer closer.Close()

	svr := favoritesrv.NewServer(new(FavoriteSrvImpl),
		server.WithTracer(prometheus.NewServerTracer(":9092", "/metrics")),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 10000000000, MaxQPS: 1000000000}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: PSM}),
		server.WithSuite(tracerSuite),
	)
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
	defer rabbitmq.Rmq.ReleaseRes()
}
