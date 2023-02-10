package client

import (
	"context"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/message/messagesrv"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"testing"
)

func TestRpcMessageChat(t *testing.T) {
	c, err := messagesrv.NewClient("bytecamp.douyin.message", client.WithHostPorts("127.0.0.1:8085"))
	if err != nil {
		fmt.Println(err)
	}
	chat, err := c.MessageChat(context.Background(), &message.DouyinMessageChatRequest{
		Token:    "",
		ToUserId: 2,
	})
	if err != nil {
		return
	}
	fmt.Println(chat)
}
