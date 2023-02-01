package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
)

type CommentActionService struct {
	ctx context.Context
}

// 创建一个评论服务
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// 评论服务实现
func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) error {
	// 根据请求创建新的评论
	if req.ActionType == 1 {
		commentModel := &db.Comment{
			VideoId: int(req.VideoId),
			UserId:  int(req.UserId),
			Content: *req.CommentText,
		}
		return db.CreateComment(s.ctx, commentModel)
	}
	// 根据请求删除评论
	if req.ActionType == 2 {
		return db.DeleteCommentById(s.ctx, int(req.VideoId), int(*req.CommentId))
	}
	// 参数不合法
	return errno.ErrBind
}
