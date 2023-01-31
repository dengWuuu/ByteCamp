/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-29 21:58:00
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-31 14:46:47
 * @FilePath: /ByteCamp/cmd/relation/handler.go
 * @Description: relation微服务handler
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"context"
	"douyin/cmd/relation/service"
	"douyin/dal/db"
	relation "douyin/kitex_gen/relation"
	"douyin/pack"
)

// RelationSrvImpl implements the last service interface defined in the IDL.
type RelationSrvImpl struct{}

var (
	relationService = service.NewRelationService(context.Background())
)

// RelationAction 登录用户对其他用户进行关注或取消关注。
func (s *RelationSrvImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	db.Init("../../config")
	//1、鉴权
	//2、入参校验
	//3、调用service层，完成关注或取消关注
	err = relationService.RelationAction(req)
	//4、返回结果
	return pack.BuildRelationActionResponse(err), err
}

// 登录用户关注的所有用户列表。
func (s *RelationSrvImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	//1、鉴权
	return
}

// 所有关注登录用户的粉丝列表。
func (s *RelationSrvImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
