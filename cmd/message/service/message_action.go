package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/message"
	"douyin/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) MessageActionService {
	return MessageActionService{ctx: ctx}
}

func (messageActionService MessageActionService) MessageAction(req *message.DouyinRelationActionRequest) error {
	fromId := middleware.GetUserIdFromTokenString(req.Token)
	message := db.Message{
		Model:      gorm.Model{},
		FromUserId: fromId,
		ToUserId:   uint(req.ToUserId),
		Content:    req.Content,
		Version:    optimisticlock.Version{},
	}
	err := db.CreateMessage(&message)
	if err != nil {
		klog.Fatal("在 message action 中创建 message 失败")
		return err
	}
	return nil
}
