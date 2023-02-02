/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-01 16:41:53
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 16:50:26
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

// RelationActionParam relation 微服务参数
type RelationActionParam struct {
	Token      string `query:"token"`
	ToUserId   int64  `query:"to_user_id"`
	ActionType int32  `query:"action_type"`
}

type FollowListParam struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type FollowerListParam struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// CommentActionParam comment操作服务输入参数
type CommentActionParam struct {
	UserId      int64   `json:"user_id,omitempty"`      // 用户id
	Token       string  `json:"token,omitempty"`        // 用户鉴权token
	VideoId     int64   `json:"video_id,omitempty"`     // 视频id
	ActionType  int32   `json:"action_type,omitempty"`  // 1-发布评论，2-删除评论
	CommentText *string `json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `json:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
}

// CommentListParam comment获取的服务输入参数
type CommentListParam struct {
	Token   string `json:"token,omitempty"`    // 用户鉴权token
	VideoId int64  `json:"video_id,omitempty"` // 视频id
}

// 点赞操作 handler 输入参数
type FavoriteActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	VideoId    int64  `json:"video_id,omitempty"`    // 视频id
	ActionType int32  `json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

// 获取点赞视频的 handler 输入参数
type FavoriteListParam struct {
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}
