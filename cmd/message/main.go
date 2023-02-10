package main

import (
	"douyin/dal"
	message "douyin/kitex_gen/message/messagesrv"
	"log"
)

// Init Relation RPC Server 端配置初始化
func Init() {
	dal.Init()
}

func main() {
	svr := message.NewServer(new(MessageSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
