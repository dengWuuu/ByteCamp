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

func TestRpcMessageAction(t *testing.T) {
	c, err := messagesrv.NewClient("bytecamp.douyin.message", client.WithHostPorts("127.0.0.1:8085"))
	if err != nil {
		fmt.Println(err)
	}
	action, err := c.MessageAction(context.Background(), &message.DouyinRelationActionRequest{
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzYxMDg1NDYsImlkZW50aXR5Ijp7IklEIjoyMywiQ3JlYXRlZEF0IjoiMjAyMy0wMS0zMVQyMjo1OTo0OS43MyswODowMCIsIlVwZGF0ZWRBdCI6IjIwMjMtMDEtMzFUMjI6NTk6NDkuNzMrMDg6MDAiLCJEZWxldGVkQXQiOm51bGwsIk5hbWUiOiJ6eTAyIiwiUGFzc3dvcmQiOiIkMmEkMTAkelI4RDFwNmtIWFAuWW8vU1BBWFVodVV6SFd1VmtsOS9xa3N5d1RoYzRSWi80MzdCNzF3OXEiLCJmb2xsb3dpbmdfY291bnQiOjAsImZvbGxvd2VyX2NvdW50IjowLCJWZXJzaW9uIjoxfSwib3JpZ19pYXQiOjE2NzYxMDQ5NDZ9.i9q242iVoMY4IA-Duj2eeqCM5IU4Jttz6v5nJ7vcKzk",
		ToUserId:   2,
		ActionType: 1,
		Content:    "asdfagagsfg",
	})
	if err != nil {
		return
	}
	fmt.Println(action)
}
