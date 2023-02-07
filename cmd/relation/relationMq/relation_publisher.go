package relationMq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

func RelationActionMqSend(message []byte) {
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

	// 2.声明交换器
	err = relationMq.Channel.ExchangeDeclare(
		relationMq.Exchange, // 交换器名
		"topic",             // exchange type：一般用fanout、direct、topic
		true,                // 是否持久化
		false,               // 是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
		false,               // 是否具有排他性
		false,               // 是否阻塞
		nil,                 // 额外属性
	)
	if err != nil {
		klog.Info("relation模块声明交换器失败")
		panic(err)
	}

	// 3.将队列绑定到交换机上
	err = relationMq.Channel.QueueBind(
		relationMq.QueueName,  // 队列名
		relationMq.RoutingKey, // 路由key
		relationMq.Exchange,   // 交换器名
		false,
		nil,
	)

	if err != nil {
		klog.Info("relation模块绑定队列失败")
		panic(err)
	}

	// 4.发送消息
	err = relationMq.Channel.Publish(
		relationMq.Exchange,   // 交换器名
		relationMq.RoutingKey, // 路由key,需要和队列的绑定key一致
		false,                 // 如果为true，根据exchange类型和routekey规则，如果无法找到符合条件的队列，那么会把发送的消息返回给发送者
		false,                 // 如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把发送的消息返回给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		klog.Info("relation模块发送消息失败")
		panic(err)
	}
}
