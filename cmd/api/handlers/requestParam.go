/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 16:41:53
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 18:40:34
 * @FilePath: /ByteCamp/cmd/api/handlers/requestParam.go
 * @Description: 用于定义handler传入参数,方便json绑定
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse pack response
func SendResponse(c *app.RequestContext, response interface{}) {
	c.JSON(consts.StatusOK, response)
}

// UserRegisterParam 用户注册handler传入参数
type UserRegisterParam struct {
	UserName string `json:"username"` // 用户名
	PassWord string `json:"password"` // 用户密码
}

// UserParam 用户输出参数
type UserParam struct {
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

//relation 微服务参数
type RelationActionParam struct {
	Token      string `json:"token"`
	ToUserId   int64  `json:"to_userid"`
	ActionType int32  `json:"action_type"`
}

type FollowListParam struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type FollowerListParam struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
