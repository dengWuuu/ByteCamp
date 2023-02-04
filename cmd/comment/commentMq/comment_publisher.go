package commentMq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

func CommentActionMqSend(message []byte) {
	_, err := commentMq.Channel.QueueDeclare(
		commentMq.QueueName,
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
		klog.Info("Add声明队列失败")
		panic(err)
	}
	// 2.声明交换器
	err = commentMq.Channel.ExchangeDeclare(
		commentMq.Exchange, //交换器名
		"topic",            //exchange type：一般用fanout、direct、topic
		true,               // 是否持久化
		false,              //是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
		false,              //设置是否内置的。true表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
		false,              // 是否阻塞
		nil,                // 额外属性
	)
	if err != nil {
		klog.Info("Add声明交换器失败", err)
		return
	}
	// 3.建立Binding(可随心所欲建立多个绑定关系)
	err = commentMq.Channel.QueueBind(
		commentMq.QueueName,  // 绑定的队列名称
		commentMq.RoutingKey, // bindkey 用于消息路由分发的key
		commentMq.Exchange,   // 绑定的exchange名
		false,                // 是否阻塞
		nil,                  // 额外属性
	)
	if err != nil {
		klog.Info("Add绑定队列和交换器失败", err)
		return
	}

	err = commentMq.Channel.Publish(
		commentMq.Exchange,
		commentMq.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		klog.Info("CommendAddMq发送消息失败")
		panic(err)
	}
}
