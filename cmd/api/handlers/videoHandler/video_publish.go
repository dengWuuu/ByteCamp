package videoHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func PublishAction(ctx context.Context, c *app.RequestContext) {
	var param handlers.VideoPublishActionParam
	err := c.Bind(&param)
	if err != nil {
		hlog.Fatalf("参数绑定失败")
		panic(err)
	}
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	var param handlers.VideoPublishListParam
	err := c.Bind(&param)
	if err != nil {
		hlog.Fatalf("参数绑定失败")
		panic(err)
	}
}
