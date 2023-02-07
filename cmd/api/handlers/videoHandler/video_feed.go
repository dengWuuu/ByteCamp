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
)

func Feed(ctx context.Context, c *app.RequestContext) {
	var param handlers.VideoFeedParam
	err := c.Bind(&param)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}

	hlog.CtxInfof(ctx, "Feed Req:%v", param)
	rpcResp, err := rpc.GetUserFeed(ctx, &video.DouyinFeedRequest{
		Token:      &param.Token,
		LatestTime: &param.LatestTime,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	hlog.CtxInfof(ctx, "Feed Resp:%v", rpcResp)
	handlers.SendResponse(c, rpcResp)
}
