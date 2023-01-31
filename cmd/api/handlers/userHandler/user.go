package userHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pack"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// Register 传递注册http请求到rpc服务
func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar handlers.UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		handlers.SendResponse(c, pack.BuildUserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Register(ctx, &user.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildUserRegisterResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, resp)
}
