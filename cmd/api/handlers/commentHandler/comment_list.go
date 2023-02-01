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

// 传送http请求上下文到rpc客户端，并且获得客户端的响应
func CommentList(ctx context.Context, c *app.RequestContext) {
	var commentListParam handlers.CommentListParam
	body, err := c.Body()
	if err != nil {
		hlog.Fatalf("获取请求体失败")
		panic(err)
	}
	err = json.Unmarshal(body, &commentListParam)
	if err != nil {
		hlog.Fatal("序列化评论获取请求参数失败")
		panic(err)
	}
	// 检查参数
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
	handlers.SendResponse(c, rpcResp)
}
