package service

import (
	"context"
	"douyin/cmd/message/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/message"
	"douyin/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) MessageChatService {
	return MessageChatService{ctx: ctx}
}

func (messageChatService MessageChatService) MessageChat(req *message.DouyinMessageChatRequest) ([]*message.Message, error) {
	fromId := middleware.GetUserIdFromTokenString(req.Token)
	chats, err := db.GetUserMessageChat(fromId, uint(req.ToUserId))
	if err != nil {
		klog.Fatal("message chats service 获取用户聊天数据失败")
		return nil, err
	}
	messages := pack.Messages(chats)

	return messages, nil
}
