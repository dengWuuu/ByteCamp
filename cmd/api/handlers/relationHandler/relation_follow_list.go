/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-06 02:18:03
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-07 14:16:16
 * @FilePath: \ByteCamp\cmd\api\handlers\relationHandler\relation_follow_list.go
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

func FollowList(ctx context.Context, c *app.RequestContext) {
	var param handlers.FollowListParam
	//1、绑定http参数
	err := c.Bind(&param)
	if err != nil {
		hlog.Infof("获取请求体失败")
		panic(err)
	}
	//2、入参校验
	if param.UserId == 0 {
		handlers.SendResponse(c, pack.BuildRelationFollowingListResp(nil, errno.ErrBind))
		return
	}

	//3、调用rpc
	resp, err := rpc.FollowList(ctx, &relation.DouyinRelationFollowListRequest{
		UserId: param.UserId,
		Token:  param.Token,
	})
	if err != nil {
		hlog.Info("调用rpc失败")
		hlog.Infof("err:%v", err.Error())
		handlers.SendResponse(c, pack.BuildRelationFollowingListResp(nil, errno.ErrBind))
		return
	}
	c.JSON(200, utils.H{
		"status_code": resp.StatusCode, // 状态码，0-成功，其他值-失败
		"status_msg":  resp.StatusMsg,  // 返回状态描述
		"user_list":   resp.UserList,
	})
}
