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
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

func GetUserById(ctx context.Context, c *app.RequestContext) {
	var userParam handlers.UserParam
	body, err := c.Body()
	if err != nil {
		hlog.Fatalf("获取请求体失败")
		panic(err)
	}
	err = json.Unmarshal(body, &userParam)
	if err != nil {
		hlog.Fatal("序列化查找请求参数失败")
		panic(err)
	}

	id, err := strconv.Atoi(strconv.FormatInt(userParam.UserId, 10))
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
	c.JSON(consts.StatusOK, utils.H{
		"status_code": resp.StatusCode, // 状态码，0-成功，其他值-失败
		"status_msg":  resp.StatusMsg,  // 返回状态描述
		"user":        resp.User,       // 用户信息
	})
}
