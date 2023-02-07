package rabbitmq

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

const MQURL = "amqp://root:Aa1076766987@81.70.207.243:5672/"

type RabbitMQ struct {
	conn       *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
	Exchange   string
	RoutingKey string
	mqurl      string
}

var Rmq *RabbitMQ

// InitRabbitMQ 初始化RabbitMQ的连接和通道。
func InitRabbitMQ() {
	Rmq = &RabbitMQ{
		mqurl: MQURL,
	}
	dial, err := amqp.Dial(Rmq.mqurl)
	if err != nil {
		Rmq.failOnErr(err, "创建连接失败")
		return
	}
	Rmq.conn = dial
}

func NewRabbitMq(queueName, exchange, routingKey string) *RabbitMQ {
	rabbitMQ := RabbitMQ{
		QueueName:  queueName,
		conn:       Rmq.conn,
		Exchange:   exchange,
		RoutingKey: routingKey,
		mqurl:      MQURL,
	}
	var err error
	// 创建Channel
	rabbitMQ.Channel, err = rabbitMQ.conn.Channel()
	if err != nil {
		Rmq.failOnErr(err, "创建channel失败")
	}
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
	err := r.conn.Close()
	if err != nil {
		Rmq.failOnErr(err, "连接关闭失败")
		return
	}
}
