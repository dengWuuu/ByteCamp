package main

import (
	"douyin/dal"
	"douyin/kitex_gen/video/videosrv"
	"douyin/pkg/nacos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
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
	viper.SetConfigName("videoService")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config")
	errV := viper.ReadInConfig()
	if errV != nil {
		hlog.Fatal("启动rpc relation 服务器时读取配置文件失败")
		return
	}

	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	//nacos
	r := registry.NewNacosRegistry(nacos.InitNacos())
	addr, _ := net.ResolveTCPAddr("tcp", viper.GetString("Server.Address")+":"+viper.GetString("Server.Port"))
	svr := videosrv.NewServer(new(VideoSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: viper.GetString("Server.Name")}))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
