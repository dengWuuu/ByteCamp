/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-04 19:52:09
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-06 02:33:50
 * @FilePath: \ByteCamp\cmd\relation\relationMq\relation_rabbitmq.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package relationMq

import (
	"douyin/pkg/rabbitmq"

	"github.com/cloudwego/kitex/pkg/klog"
)

var relationMq *rabbitmq.RabbitMQ

func InitRelationMq() {
	rabbitmq.InitRabbitMQ()
	klog.Info("relation模块初始化rabbitmq连接成功")
	relationMq = rabbitmq.NewRabbitMq("relation_queue", "relation_exchange", "relation")
	klog.Info("relation模块初始化rabbitmq channel成功")
	go relationConsumer()
}
