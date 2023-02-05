/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-04 20:32:01
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-04 20:37:32
 * @FilePath: /ByteCamp/cmd/relation/relationMq/relation_consumer.go
 * @Description: mq消费者代码
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package relationMq

import "github.com/cloudwego/kitex/pkg/klog"

func RelationConsumer() {
	//1.声明队列
	_, err := relationMq.Channel.QueueDeclare(
		relationMq.QueueName,
		//是否持久化
		true,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		klog.Info("relation模块声明队列失败")
		panic(err)
	}

	//2.接收消息
	msgChannel, err := relationMq.Channel.Consume(
		relationMq.QueueName,
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		klog.Info("relation模块接收消息失败")
	}
	//3.处理消息
	for msg := range msgChannel {
		klog.Info("relation模块接收到消息:", string(msg.Body))
		err := msg.Ack(true)
		if err != nil {
			klog.Info("ack失败")
			return
		}
	}
}
