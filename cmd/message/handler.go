package main

import (
	"context"
	"douyin/cmd/message/pack"
	"douyin/cmd/message/service"
	"douyin/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"

	message "douyin/kitex_gen/message"
)

// MessageSrvImpl implements the last service interface defined in the IDL.
type MessageSrvImpl struct{}

// MessageChat implements the MessageSrvImpl interface.
func (s *MessageSrvImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	if req.ToUserId < 0 {
		resp = pack.BuildMessageChatResp(errno.ErrBind)
		return resp, nil
	}

	rpcMessage, err := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		klog.Fatal("MessageChat handler 获取 messages失败")
		return pack.BuildMessageChatResp(err), nil
	}
	resp = pack.BuildMessageChatResp(errno.Success)
	resp.MessageList = rpcMessage
	return resp, nil
}

// MessageAction implements the MessageSrvImpl interface.
func (s *MessageSrvImpl) MessageAction(ctx context.Context, req *message.DouyinRelationActionRequest) (resp *message.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}
