/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-04 20:32:01
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-05 11:08:06
 * @FilePath: /ByteCamp/cmd/relation/relationMq/relation_consumer.go
 * @Description: mq消费者代码
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package relationMq

import (
	"encoding/json"
	"errors"

	"douyin/dal/db"
	"douyin/kitex_gen/relation"

	"github.com/cloudwego/kitex/pkg/klog"
)

func relationConsumer() {
	// 1.声明队列
	_, err := relationMq.Channel.QueueDeclare(
		relationMq.QueueName,
		// 是否持久化
		true,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		klog.Info("relation模块声明队列失败")
		panic(err)
	}

	// 2.接收消息
	msgChannel, err := relationMq.Channel.Consume(
		relationMq.QueueName,
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		klog.Info("relation模块接收消息失败")
	}
	// 3.处理消息
	for msg := range msgChannel {
		klog.Info("relation模块接收到消息:", string(msg.Body))
		// TODO:根据消息内容进行处理
		var req *relation.DouyinRelationActionRequest
		err := json.Unmarshal(msg.Body, &req)
		if err != nil {
			klog.Fatalf("relation模块mq解析消息失败: %v", err)
		}
		// err = msg.Ack(true)
		// if err != nil {
		// 	klog.Info("ack失败")
		// 	return
		// }
		err = relationActionHandle(req)
		if err != nil {
			klog.Infof("relation模块处理消息失败：%v", err)
		}
	}
}

func relationActionHandle(req *relation.DouyinRelationActionRequest) error {
	if req.ActionType == 1 {
		// 首先判断之前是否已经关注过
		// 如果已经关注过，直接将其cancel字段重置为0即可
		// 如果没有关注过，直接插入一条新的关注记录
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
