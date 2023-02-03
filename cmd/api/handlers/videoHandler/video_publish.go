package videoHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/pack"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
)

func PublishAction(ctx context.Context, c *app.RequestContext) {
	var param handlers.VideoPublishActionParam
	err := c.Bind(&param)
	if err != nil {
		hlog.Fatalf("参数绑定失败")
		panic(err)
	}

	file, err := param.Data.Open()
	fileData, err := io.ReadAll(file)
	if err != nil {
		return
	}
	rpcResp, err := rpc.PublishAction(ctx, &video.DouyinPublishActionRequest{
		Token: param.Token,
		Title: param.Title,
		Data:  fileData,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, rpcResp)
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	var param handlers.VideoPublishListParam
	err := c.Bind(&param)
	if err != nil {
		hlog.Fatalf("参数绑定失败")
		panic(err)
	}

	rpcResp, err := rpc.PublishList(ctx, &video.DouyinPublishListRequest{
		Token:  param.Token,
		UserId: param.UserId,
	})

	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, rpcResp)
}
