/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-04 19:52:09
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-04 19:54:37
 * @FilePath: /ByteCamp/cmd/relation/relationMq/relation_rabbitmq.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package relationMq

import "douyin/pkg/rabbitmq"

var relationMq *rabbitmq.RabbitMQ

func InitCommentMq() {
	rabbitmq.InitRabbitMQ()
	relationMq = rabbitmq.NewRabbitMq("comment_queue", "comment_exchange", "comment")
}
