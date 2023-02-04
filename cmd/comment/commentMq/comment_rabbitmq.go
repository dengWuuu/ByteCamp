package commentMq

import (
	"douyin/pkg/rabbitmq"
)

var commentMq *rabbitmq.RabbitMQ

func InitCommentMq() {
	rabbitmq.InitRabbitMQ()
	commentMq = rabbitmq.NewRabbitMq("comment_queue", "comment_exchange", "comment")
}
