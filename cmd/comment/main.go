package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"douyin/cmd/comment/commentMq"
	"douyin/cmd/comment/service"
	"douyin/dal"
	comment "douyin/kitex_gen/comment/commentsrv"
	dfa "douyin/pkg/dfa"
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
	commentMq.InitCommentMq()
	go commentMq.CommentConsumer()
}

func main() {
	Init()
	PSM := "bytecamp.douyin.comment"
	Address := "127.0.0.1"
	Port := 8083

	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", Address, Port))
	// nacos
	r := registry.NewNacosRegistry(nacos.InitNacos())

	// init sensitive words check
	service.FDA = dfa.New()
	sensitiveFile, err := os.Open("cmd/comment/config/sensitive.txt")
	if err != nil {
		panic("读取敏感词文件失败")
	}
	defer sensitiveFile.Close()
	reader := bufio.NewReader(sensitiveFile)
	fileContent, err := reader.ReadString('\n')
	if err == io.EOF {
		klog.Info("读取敏感词文件成功")
	}
	sensitiveWords := strings.Fields(fileContent)
	service.FDA.AddBadWords(sensitiveWords)

	//jaeger
	//tracerSuite, closer := jaeger.InitJaegerServer("comment-server")
	//defer closer.Close()

	svr := comment.NewServer(new(CommentSrvImpl),
		server.WithTracer(prometheus.NewServerTracer(":9091", "/metrics")),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000000000, MaxQPS: 1000000000}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: PSM}),
		//server.WithSuite(tracerSuite),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
	defer rabbitmq.Rmq.ReleaseRes()
}
