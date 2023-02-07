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

func GetUserById(ctx context.Context, c *app.RequestContext) {
	var userParam handlers.UserParam
	err := c.Bind(&userParam)
	if err != nil {
		return
	}

	if userParam.UserId < 0 {
		handlers.SendResponse(c, pack.BuildGetUserResp(errno.ErrBind))
		return
	}
	resp, err := rpc.GetUserById(ctx, &user.DouyinUserRequest{
		UserId: userParam.UserId,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildGetUserResp(errno.ConvertErr(err)))
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": resp.StatusCode, // 状态码，0-成功，其他值-失败
		"status_msg":  resp.StatusMsg,  // 返回状态描述
		"user":        resp.User,       // 用户信息
	})
}
