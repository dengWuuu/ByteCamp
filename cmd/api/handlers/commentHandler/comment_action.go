package commentHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/comment/pack"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// CommentAction 传递请求上下文到comment服务的rpc客户端并获取对应的响应
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var commentActionParam handlers.CommentActionParam
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	// 检查参数
	vid, err := strconv.Atoi(video_id)
	if err != nil {
		handlers.SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		handlers.SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}

	commentActionParam.Token = token
	commentActionParam.VideoId = int64(vid)
	commentActionParam.ActionType = int32(act)

	if commentActionParam.ActionType != 1 && commentActionParam.ActionType != 2 {
		handlers.SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}
	// 封装请求传送到rpc客户端
	queryParam := &comment.DouyinCommentActionRequest{
		Token:      commentActionParam.Token,
		VideoId:    commentActionParam.VideoId,
		ActionType: commentActionParam.ActionType,
	}
	if commentActionParam.ActionType == 1 {
		comment_text := c.Query("comment_text")
		queryParam.CommentText = &comment_text
	}
	if commentActionParam.ActionType == 2 {
		comment_id := c.Query("comment_id")
		cid, err := strconv.Atoi(comment_id)
		if err != nil {
			handlers.SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
			return
		}
		cid64 := int64(cid)
		queryParam.CommentId = &cid64
	}
	rpcResp, err := rpc.CommentAction(ctx, queryParam)
	if err != nil {
		handlers.SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, rpcResp)
}
