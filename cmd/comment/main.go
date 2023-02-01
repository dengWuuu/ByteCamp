package main

import (
	"douyin/dal"
	comment "douyin/kitex_gen/comment/commentsrv"
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
	viper.SetConfigName("commentService")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config")
	errV := viper.ReadInConfig()
	if errV != nil {
		hlog.Fatal("启动rpc comment服务器时读取配置文件失败")
		return
	}
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", viper.GetString("Server.Address")+":"+viper.GetString("Server.Port"))
	svr := comment.NewServer(new(CommentSrvImpl),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: viper.GetString("Server.Name")}))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
