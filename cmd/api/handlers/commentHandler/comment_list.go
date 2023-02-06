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
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// CommentList 传送http请求上下文到rpc客户端，并且获得客户端的响应
func CommentList(ctx context.Context, c *app.RequestContext) {
	var commentListParam handlers.CommentListParam
	// 获取参数
	token := c.Query("token")
	video_id := c.Query("video_id")
	// 检查参数
	vid, err := strconv.Atoi(video_id)
	if err != nil {
		handlers.SendResponse(c, pack.BuildCommentListResp(errno.ErrBind))
		return
	}
	commentListParam.Token = token
	commentListParam.VideoId = int64(vid)
	if commentListParam.VideoId <= 0 {
		handlers.SendResponse(c, pack.BuildCommentListResp(errno.ErrBind))
		return
	}
	// 封装请求到rpc客户端
	queryParam := &comment.DouyinCommentListRequest{
		VideoId: commentListParam.VideoId,
		Token:   commentListParam.Token,
	}
	rpcResp, err := rpc.CommentList(ctx, queryParam)
	if err != nil {
		handlers.SendResponse(c, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}
	// 修改返回响应的格式
	c.JSON(200, utils.H{
		"status_code":  rpcResp.StatusCode,
		"status_msg":   rpcResp.StatusMsg,
		"comment_list": rpcResp.CommentList,
	})
}
