package rpc

import (
	"douyin/kitex_gen/comment/commentsrv"
	"os"
)

var commentClient commentsrv.Client

// comment客户端初始化
func InitCommentRpc() {
	// 读取配置
	path, err1 := os.Getwd()
	if err1 != nil {
		// 语言宕机
		panic(err1)
	}
}
