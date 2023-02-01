package service

import (
	"context"
	"douyin/cmd/comment/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest) ([]*comment.Comment, error) {
	res, err := db.GetCommentByVideoId(s.ctx, req.VideoId)
	// 判断错误
	if err != nil {
		return nil, err
	}
	// 打包成rpc通信使用的数据结构体
	comments, err := pack.Comments(s.ctx, res)
	// 判断错误
	if err != nil {
		return nil, err
	}
	return comments, nil
}
