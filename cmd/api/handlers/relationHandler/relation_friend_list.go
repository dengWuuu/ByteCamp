package relationHandler

import (
	"context"

	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/relation/pack"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func FriendList(ctx context.Context, c *app.RequestContext) {
	var param handlers.FriendListParam
	// 1、绑定http参数
	err := c.Bind(&param)
	if err != nil {
		hlog.Infof("序列化朋友列表请求参数失败")
		panic(err)
	}
	// 2、入参校验
	if param.UserId == 0 {
		resp := pack.BuildRelationFollowerListResp(nil, errno.ErrBind)
		c.JSON(200, utils.H{
			"status_code": resp.StatusCode, // 状态码，0-成功，其他值-失败
			"status_msg":  resp.StatusMsg,  // 返回状态描述
			"user_list":   nil,
		})
	}

	// 3、调用rpc
	resp, err := rpc.FriendList(ctx, &relation.DouyinRelationFriendListRequest{
		UserId: param.UserId,
		Token:  param.Token,
	})
	if err != nil {
		resp := pack.BuildRelationFollowerListResp(nil, err)
		c.JSON(200, utils.H{
			"status_code": resp.StatusCode, // 状态码，0-成功，其他值-失败
			"status_msg":  resp.StatusMsg,  // 返回状态描述
			"user_list":   nil,
		})
	}
	c.JSON(200, utils.H{
		"status_code": resp.StatusCode, // 状态码，0-成功，其他值-失败
		"status_msg":  resp.StatusMsg,  // 返回状态描述
		"user_list":   resp.UserList,
	})
}
