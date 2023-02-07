package pack

import (
	"errors"

	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
)

// 根据错误信息包装RPC返回数据结构体
func BuildCommentActionResp(err error) *comment.DouyinCommentActionResponse {
	if err == nil {
		return commentActionResp(errno.Success)
	}
	e := errno.ErrNo{}
	// 查看是否在错误链上能够查询到
	if errors.As(err, &e) {
		return commentActionResp(e)
	}
	// 未知错误，提出错误信息包装为新的未知错误
	e = errno.ErrUnknown.WithMessage(err.Error())
	return commentActionResp(e)
}

func BuildCommentListResp(err error) *comment.DouyinCommentListResponse {
	if err == nil {
		return commentListResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentListResp(e)
	}
	e = errno.ErrUnknown.WithMessage(err.Error())
	return commentListResp(e)
}

// 自定义的错误来封装响应结果
func commentActionResp(err errno.ErrNo) *comment.DouyinCommentActionResponse {
	return &comment.DouyinCommentActionResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}

func commentListResp(err errno.ErrNo) *comment.DouyinCommentListResponse {
	return &comment.DouyinCommentListResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}
