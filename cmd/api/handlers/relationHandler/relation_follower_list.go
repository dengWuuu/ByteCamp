/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-02 16:44:03
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 17:14:24
 * @FilePath: /ByteCamp/cmd/api/handlers/relationHandler/relation_follower_list.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
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

func FollowerList(ctx context.Context, c *app.RequestContext) {
	var param handlers.FollowerListParam
	// 1、绑定http参数
	err := c.Bind(&param)
	if err != nil {
		hlog.Fatal("序列化粉丝列表请求参数失败")
		panic(err)
	}
	// 2、入参校验
	if param.UserId == 0 {
		handlers.SendResponse(c, pack.BuildRelationFollowerListResp(nil, errno.ErrBind))
		return
	}

	// 3、调用rpc
	resp, err := rpc.FollowerList(ctx, &relation.DouyinRelationFollowerListRequest{
		UserId: param.UserId,
		Token:  param.Token,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildRelationFollowerListResp(nil, errno.ErrBind))
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.StatusCode, // 状态码，0-成功，其他值-失败
		"status_msg":  resp.StatusMsg,  // 返回状态描述
		"user_list":   resp.UserList,
	})
}
