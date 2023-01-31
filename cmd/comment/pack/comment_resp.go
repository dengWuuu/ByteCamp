package pack

import (
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
	"errors"
)

func BuildBaseCommentResp(m *comment.Comment, err error) *comment.DouyinCommentActionResponse {
	if err == nil {
		return baseSuccessResp(m, errno.Success)
	}
	e := errno.ErrNo{}
	//如果是定义的错误则打印
	if errors.As(err, &e) {
		return baseFailResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return baseFailResp(s)
}

func baseSuccessResp(m *comment.Comment, err errno.ErrNo) *comment.DouyinCommentActionResponse {
	return &comment.DouyinCommentActionResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
		Comment:    m, // 没有评论内容
	}
}
func baseFailResp(err errno.ErrNo) *comment.DouyinCommentActionResponse {
	return &comment.DouyinCommentActionResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
		Comment:    nil,
	}
}
