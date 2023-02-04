package main

import (
	"douyin/cmd/comment/commentMq"
	"douyin/dal"
	comment "douyin/kitex_gen/comment/commentsrv"
	"douyin/pkg/nacos"
	"douyin/pkg/rabbitmq"
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
	commentMq.InitCommentMq()
	commentMq.CommentConsumer()
}
func main() {
	Init()
	PSM := "bytecamp.douyin.comment"
	Address := "127.0.0.1"
	Port := 8083
	//Port, err := nacos.GetFreePort()
	//if err != nil{
	//	panic(err)
	//}
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", Address, Port))
	//nacos
	r := registry.NewNacosRegistry(nacos.InitNacos())
	svr := comment.NewServer(new(CommentSrvImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: PSM}))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
	defer rabbitmq.Rmq.ReleaseRes()
}
