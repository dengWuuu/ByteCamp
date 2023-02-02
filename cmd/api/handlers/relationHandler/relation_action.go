/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 14:46:43
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 17:02:15
 * @FilePath: /ByteCamp/cmd/api/handlers/relationHandler/relation_action.go
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
	"douyin/pkg/middleware"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func RelationAction(ctx context.Context, c *app.RequestContext) {
	//TODO:修改参数请求方式

	//1、绑定http参数
	var param handlers.RelationActionParam
	err := c.Bind(&param)
	if err != nil {
		hlog.Fatalf("参数绑定失败")
		panic(err)
	}

	//2、入参校验
	if param.ToUserId == 0 || param.ActionType == 0 {
		handlers.SendResponse(c, pack.BuildRelationActionResponse(errno.ErrBind))
		return
	}

	//3、调用rpc
	//TODO:从token中解析出userId,这里暂时先将userId写死为1(Done)
	userId := middleware.GetUserIdFromJwtToken(ctx, c)
	resp, err := rpc.RelationAction(ctx, &relation.DouyinRelationActionRequest{
		UserId:     int64(userId),
		Token:      param.Token,
		ToUserId:   param.ToUserId,
		ActionType: param.ActionType,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildRelationActionResponse(errno.ConvertErr(err)))
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.StatusCode, // 状态码，0-成功，其他值-失败
		"status_msg":  resp.StatusMsg,  // 返回状态描述
	})
}
