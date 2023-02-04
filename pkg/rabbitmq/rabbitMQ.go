package rabbitmq

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

const MQURL = "amqp://root:Aa1076766987@81.70.207.243:5672/"

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// routing Key
	RoutingKey string
	//MQ链接字符串
	Mqurl string
}

var Rmq *RabbitMQ

// InitRabbitMQ 初始化RabbitMQ的连接和通道。
func InitRabbitMQ() {
	Rmq = &RabbitMQ{
		Mqurl: MQURL,
	}
	dial, err := amqp.Dial(Rmq.Mqurl)
	if err != nil {
		Rmq.failOnErr(err, "创建连接失败")
		return
	}
	Rmq.Conn = dial
}
func NewRabbitMq(queueName, exchange, routingKey string) *RabbitMQ {
	rabbitMQ := RabbitMQ{
		QueueName:  queueName,
		Exchange:   exchange,
		RoutingKey: routingKey,
		Mqurl:      MQURL,
	}
	var err error
	//创建rabbitmq连接
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)
	Rmq.failOnErr(err, "创建连接失败")

	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	Rmq.failOnErr(err, "创建channel失败")
	return &rabbitMQ
}

// 连接出错时，输出错误信息。
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		klog.Fatalf("%s:%s\n", err, message)
		panic(fmt.Sprintf("%s:%s\n", err, message))
	}
}

// ReleaseRes 关闭mq通道和mq的连接。
func (r *RabbitMQ) ReleaseRes() {
	err := r.Conn.Close()
	if err != nil {
		Rmq.failOnErr(err, "连接关闭失败")
		return
	}
	err = r.Channel.Close()
	if err != nil {
		Rmq.failOnErr(err, "信道关闭失败")
		return
	}
}
