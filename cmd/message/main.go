package main

import (
	message "douyin/kitex_gen/message/messagesrv"
	"log"
)

func main() {
	svr := message.NewServer(new(MessageSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
