package userHandler

import (
	"context"

	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// Register 传递注册http请求到rpc服务
func Register(ctx context.Context, c *app.RequestContext) {
	var registerParam handlers.UserRegisterParam
	err := c.Bind(&registerParam)
	if err != nil {
		return
	}

	if len(registerParam.UserName) == 0 || len(registerParam.PassWord) == 0 {
		handlers.SendResponse(c, pack.BuildUserRegisterResp(errno.ErrBind))
		return
	}

	_, err = rpc.Register(ctx, &user.DouyinUserRegisterRequest{
		Username: registerParam.UserName,
		Password: registerParam.PassWord,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildUserRegisterResp(errno.ConvertErr(err)))
		return
	}
	//login after register
	Login(ctx, c)
}
