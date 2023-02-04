package videoHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/pack"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"
	"douyin/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"strconv"
)

func PublishAction(ctx context.Context, c *app.RequestContext) {
	var param handlers.VideoPublishActionParam
	err := c.Bind(&param)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}

	hlog.CtxInfof(ctx, "Publish Req:%v", param)
	file, err := param.Data.Open()
	fileData, err := io.ReadAll(file)
	userID := middleware.GetUserIdFromJwtToken(ctx, c)
	if err != nil {
		return
	}
	rpcResp, err := rpc.PublishAction(ctx, &video.DouyinPublishActionRequest{
		Token: strconv.Itoa(int(userID)),
		Title: param.Title,
		Data:  fileData,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	hlog.CtxInfof(ctx, "Publish Resp:%v", rpcResp)
	handlers.SendResponse(c, rpcResp)
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	var param handlers.VideoPublishListParam
	err := c.Bind(&param)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}

	hlog.CtxInfof(ctx, "PublishList Req:%v", param)
	rpcResp, err := rpc.PublishList(ctx, &video.DouyinPublishListRequest{
		Token:  param.Token,
		UserId: param.UserId,
	})

	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	hlog.CtxInfof(ctx, "PublishList Resp:%v", rpcResp)
	handlers.SendResponse(c, rpcResp)
}
