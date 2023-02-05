/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 12:15:33
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-05 09:09:55
 * @FilePath: /ByteCamp/cmd/relation/service/relation_action.go
 * @Description: relationAction接口对应service
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import (
	"context"
	"douyin/cmd/relation/relationMq"
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"encoding/json"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
)

//TODO:采用redis维护用户的关注、粉丝、朋友列表
func (service *RelationService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	if req.ActionType == 1 {
		//首先判断之前是否已经关注过
		//如果已经关注过，直接将其cancel字段重置为0即可
		//如果没有关注过，直接插入一条新的关注记录
		follow, _ := db.GetFollowByUserAndTarget(req.UserId, req.ToUserId)
		if follow.ID != 0 {
			err := db.UpdateFollow(req.UserId, req.ToUserId, int(req.ActionType))
			if err != nil {
				return err
			}
		} else {
			err := db.AddRelation(int(req.UserId), int(req.ToUserId))
			if err != nil {
				return err
			}
		}
	} else if req.ActionType == 2 {
		err := db.UpdateFollow(req.UserId, req.ToUserId, int(req.ActionType))
		if err != nil {
			return err
		}
	} else {
		return errors.New("action_type error")
	}
	return nil
}

func (service *RelationService) RelationActionByRedis(req *relation.DouyinRelationActionRequest) error {
	ctx := context.Background()
	//发送mq
	msg, err := json.Marshal(req)
	if err != nil {
		klog.Info("mq msg json marshal error")
		return err
	}
	relationMq.RelationActionMqSend([]byte(msg))
	if req.ActionType == 1 {
		//更新redis
		addRedisFollowList(ctx, req.UserId, req.ToUserId)
		addRedisFollowerList(ctx, req.UserId, req.ToUserId)
		addRedisFriendsList(ctx, req.UserId, req.ToUserId)
	} else if req.ActionType == 2 {
		//更新redis
		remRedisFollowList(ctx, req.UserId, req.ToUserId)
		remRedisFollowerList(ctx, req.UserId, req.ToUserId)
		remRedisFriendsList(ctx, req.UserId, req.ToUserId)
	} else {
		return errors.New("action_type error")
	}
	return nil
}
