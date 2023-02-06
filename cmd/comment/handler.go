package main

import (
	"context"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/service"
	comment "douyin/kitex_gen/comment"
	"douyin/pkg/errno"
)

// CommentSrvImpl implements the last service interface defined in the IDL.
type CommentSrvImpl struct{}

// CommentAction implements the CommentSrvImpl interface.
// 实现RPC中的关于评论的功能接口
func (s *CommentSrvImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	commentSrv := service.NewCommentActionService(ctx)
	// 检查参数
	if req.UserId <= 0 || req.VideoId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		return pack.BuildCommentActionResp(errno.ErrBind), nil
	}
	comment, err := commentSrv.CommentAction(req)
	if err != nil {
		return pack.BuildCommentActionResp(err), nil
	}
	resp = pack.BuildCommentActionResp(err)
	resp.Comment = comment
	return resp, nil
}

// CommentList implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	commentSrv := service.NewCommentListService(ctx)
	// 检查参数
	if req.VideoId <= 0 {
		return pack.BuildCommentListResp(errno.ErrBind), nil
	}
	res, err := commentSrv.CommentList(req)
	// 出现错误
	if err != nil {
		return pack.BuildCommentListResp(err), nil
	}
	pack_res := pack.BuildCommentListResp(err)
	pack_res.CommentList = res
	return pack_res, nil
}
