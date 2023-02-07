package userHandler

import (
	"context"

	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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

	resp, err := rpc.Register(ctx, &user.DouyinUserRegisterRequest{
		Username: registerParam.UserName,
		Password: registerParam.PassWord,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildUserRegisterResp(errno.ConvertErr(err)))
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserId,
		"token":       resp.Token,
	})
}
