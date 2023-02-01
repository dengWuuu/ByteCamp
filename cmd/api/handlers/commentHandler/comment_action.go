package commentHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/comment/pack"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// 传递请求上下文到comment服务的rpc客户端并获取对应的响应
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var commentActionParam handlers.CommentActionParam
	body, err := c.Body()
	if err != nil {
		hlog.Fatalf("获取请求体失败")
		panic(err)
	}
	err = json.Unmarshal(body, &commentActionParam)
	if err != nil {
		hlog.Fatal("序列化用户评论请求参数失败")
		panic(err)
	}
	// 检查参数
	if commentActionParam.UserId <= 0 || commentActionParam.VideoId <= 0 {
		handlers.SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}
	if commentActionParam.ActionType != 1 && commentActionParam.ActionType != 2 {
		handlers.SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}
	// 封装请求传送到rpc客户端
	queryParam := &comment.DouyinCommentActionRequest{
		UserId:     commentActionParam.UserId,
		Token:      commentActionParam.Token,
		VideoId:    commentActionParam.VideoId,
		ActionType: commentActionParam.ActionType,
	}
	if commentActionParam.ActionType == 1 {
		queryParam.CommentText = commentActionParam.CommentText
	}
	if commentActionParam.ActionType == 2 {
		queryParam.CommentId = commentActionParam.CommentId
	}
	rpcResp, err := rpc.CommentAction(ctx, queryParam)
	if err != nil {
		handlers.SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, rpcResp)
}
