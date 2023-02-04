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
		hlog.Fatalf("参数绑定失败")
		panic(err)
	}

	rpcResp, err := rpc.GetUserFeed(ctx, &video.DouyinFeedRequest{
		Token:      &param.Token,
		LatestTime: &param.LatestTime,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, rpcResp)
}
