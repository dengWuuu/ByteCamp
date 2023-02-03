package videoHandler

import (
	"context"
	"douyin/cmd/api/handlers"
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

	if err != nil {
		hlog.Fatalf("参数绑定失败")
		panic(err)
	}
}
