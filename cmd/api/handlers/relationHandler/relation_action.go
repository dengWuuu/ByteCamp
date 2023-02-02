/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 14:46:43
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 23:25:09
 * @FilePath: /ByteCamp/cmd/api/handlers/relationHandler/relation_action.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package relationhandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/relation/pack"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
	"douyin/pkg/middleware"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func RelationAction(ctx context.Context, c *app.RequestContext) {
	var param handlers.RelationActionParam
	//1、绑定http参数
	body, err := c.Body()
	if err != nil {
		hlog.Fatalf("获取请求体失败")
		panic(err)
	}
	err = json.Unmarshal(body, &param)
	if err != nil {
		hlog.Fatal("序列化用户关注请求参数失败")
		panic(err)
	}
	//2、入参校验
	if param.ToUserId == 0 || param.ActionType == 0 {
		handlers.SendResponse(c, pack.BuildRelationActionResponse(errno.ErrBind))
		return
	}

	//3、调用rpc
	//TODO:从token中解析出userId,这里暂时先将userId写死为1
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
	handlers.SendResponse(c, resp)
}
