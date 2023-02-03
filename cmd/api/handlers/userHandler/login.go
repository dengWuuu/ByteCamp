package userHandler

import (
	"context"
	"douyin/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func Login(ctx context.Context, c *app.RequestContext) {
	middleware.JwtMiddleware.LoginHandler(ctx, c)
}
