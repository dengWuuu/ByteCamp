package messageHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/message"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func MessageChat(ctx context.Context, c *app.RequestContext) {

	// 1、绑定http参数
	var param handlers.MessageChatParam
	err := c.Bind(&param)
	if err != nil {
		hlog.Infof("参数绑定失败")
		panic(err)
	}

	// 2、入参校验
	if param.ToUserId < 0 {
		handlers.SendResponse(c, pack.BuildGetUserResp(errno.ErrBind))
	}
	// 3、调用rpc
	resp, err := rpc.MessageChat(ctx, &message.DouyinMessageChatRequest{
		Token:    param.Token,
		ToUserId: param.ToUserId,
	})
	if err != nil {
		hlog.Fatalf("调用Message rpc 调用失败")
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code":  resp.StatusCode,  // 状态码，0-成功，其他值-失败
		"status_msg":   resp.StatusMsg,   // 返回状态描述
		"message_list": resp.MessageList, // 聊天记录
	})

}
