package pack

import (
	"douyin/kitex_gen/message"
	"douyin/pkg/errno"
	"github.com/pkg/errors"
)

func BuildMessageChatResp(err error) *message.DouyinMessageChatResponse {
	if err == nil {
		return getMessageChatResp(errno.Success)
	}
	// 如果是定义的错误则打印
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getMessageChatResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return getMessageChatResp(s)
}
func getMessageChatResp(err errno.ErrNo) *message.DouyinMessageChatResponse {
	return &message.DouyinMessageChatResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}

func BuildMessageActionResp(err error) *message.DouyinRelationActionResponse {
	if err == nil {
		return getMessageActionResp(errno.Success)
	}
	// 如果是定义的错误则打印
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getMessageActionResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return getMessageActionResp(s)
}
func getMessageActionResp(err errno.ErrNo) *message.DouyinRelationActionResponse {
	return &message.DouyinRelationActionResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}
