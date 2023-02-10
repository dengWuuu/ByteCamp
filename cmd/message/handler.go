package main

import (
	"context"
	message "douyin/kitex_gen/message"
)

// MessageSrvImpl implements the last service interface defined in the IDL.
type MessageSrvImpl struct{}

// MessageChat implements the MessageSrvImpl interface.
func (s *MessageSrvImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageAction implements the MessageSrvImpl interface.
func (s *MessageSrvImpl) MessageAction(ctx context.Context, req *message.DouyinRelationActionRequest) (resp *message.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}
