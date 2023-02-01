package userHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/middleware/JwtUtils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
)

func GetUserById(ctx context.Context, c *app.RequestContext) {
	hlog.Infof("当前用户是?", JwtUtils.GetUserIdFromJwtToken(ctx, c))
	str := c.Query("user_id")

	id, err := strconv.Atoi(str)
	if err != nil {
		handlers.SendResponse(c, pack.BuildGetUserResp(errno.ErrBind))
		return
	}
	userId := int64(id)
	if userId < 0 {
		handlers.SendResponse(c, pack.BuildGetUserResp(errno.ErrBind))
		return
	}
	resp, err := rpc.GetUserById(ctx, &user.DouyinUserRequest{
		UserId: userId,
	})

	if err != nil {
		handlers.SendResponse(c, pack.BuildGetUserResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, resp)

}
