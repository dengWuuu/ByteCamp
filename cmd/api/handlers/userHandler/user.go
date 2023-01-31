package userHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// Register 传递注册http请求到rpc服务
func Register(ctx context.Context, c *app.RequestContext) {
	var registerParam handlers.UserRegisterParam
	body, err := c.Body()
	if err != nil {
		hlog.Fatalf("获取请求体失败")
		panic(err)
	}
	err = json.Unmarshal(body, &registerParam)
	if err != nil {
		hlog.Fatal("序列化用户注册请求参数失败")
		panic(err)
	}
	if len(registerParam.UserName) == 0 || len(registerParam.PassWord) == 0 {
		handlers.SendResponse(c, pack.BuildUserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Register(ctx, &user.DouyinUserRegisterRequest{
		Username: registerParam.UserName,
		Password: registerParam.PassWord,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildUserRegisterResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, resp)
}
